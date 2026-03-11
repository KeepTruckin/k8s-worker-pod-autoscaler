package controller

import (
	"context"
	"testing"

	wpaapi "github.com/practo/k8s-worker-pod-autoscaler/pkg/apis/workerpodautoscalermultiqueue/v1"
	fakecustomclientset "github.com/practo/k8s-worker-pod-autoscaler/pkg/generated/clientset/versioned/fake"
	wpalisters "github.com/practo/k8s-worker-pod-autoscaler/pkg/generated/listers/workerpodautoscalermultiqueue/v1"
	"github.com/practo/k8s-worker-pod-autoscaler/pkg/queue"
	"github.com/rs/zerolog"
	statsig "github.com/statsig-io/go-sdk"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appslisters "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
)

func TestSyncHandlerScalesDeploymentWithoutEnvNameMatch(t *testing.T) {
	statsig.InitializeWithOptions("secret-local", &statsig.Options{LocalMode: true})
	defer statsig.Shutdown()

	testCases := []struct {
		name           string
		env            string
		deploymentName string
	}{
		{
			name:           "deployment name does not contain env",
			env:            "development",
			deploymentName: "payments-worker",
		},
		{
			name:           "deployment name contains env",
			env:            "development",
			deploymentName: "payments-development-worker",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			namespace := "default"
			wpaName := "test-wpa"
			minReplicas := int32(1)
			maxReplicas := int32(10)
			replicas := int32(1)
			availableReplicas := int32(2)

			deployment := &appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      tc.deploymentName,
					Namespace: namespace,
				},
				Spec: appsv1.DeploymentSpec{
					Replicas: &replicas,
				},
				Status: appsv1.DeploymentStatus{
					AvailableReplicas: availableReplicas,
				},
			}

			wpa := &wpaapi.WorkerPodAutoScalerMultiQueue{
				ObjectMeta: metav1.ObjectMeta{
					Name:      wpaName,
					Namespace: namespace,
				},
				Spec: wpaapi.WorkerPodAutoScalerMultiQueueSpec{
					MinReplicas:    &minReplicas,
					MaxReplicas:    &maxReplicas,
					DeploymentName: tc.deploymentName,
				},
			}

			deploymentIndexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{})
			if err := deploymentIndexer.Add(deployment); err != nil {
				t.Fatalf("failed to add deployment to indexer: %v", err)
			}
			wpaIndexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{})
			if err := wpaIndexer.Add(wpa); err != nil {
				t.Fatalf("failed to add wpa to indexer: %v", err)
			}

			queues := queue.NewQueues(zerolog.Nop())
			stopCh := make(chan struct{})
			go queues.Sync(stopCh)
			defer close(stopCh)

			controller := &Controller{
				ctx:                        ctx,
				logger:                     zerolog.Nop(),
				customclientset:            fakecustomclientset.NewSimpleClientset(wpa.DeepCopy()),
				deploymentLister:           appslisters.NewDeploymentLister(deploymentIndexer),
				workerPodAutoScalersLister: wpalisters.NewWorkerPodAutoScalerMultiQueueLister(wpaIndexer),
				defaultMaxDisruption:       "100%",
				scaleDownDelay:             0,
				Queues:                     queues,
				env:                        tc.env,
			}

			err := controller.syncHandler(ctx, WokerPodAutoScalerEvent{
				key:  namespace + "/" + wpaName,
				name: WokerPodAutoScalerEventAdd,
			})
			if err != nil {
				t.Fatalf("syncHandler returned error: %v", err)
			}

			updatedWPA, err := controller.customclientset.K8sV1().
				WorkerPodAutoScalerMultiQueues(namespace).
				Get(ctx, wpaName, metav1.GetOptions{})
			if err != nil {
				t.Fatalf("failed to fetch updated wpa: %v", err)
			}

			if updatedWPA.Status.CurrentReplicas != replicas {
				t.Fatalf("expected status.currentReplicas=%d, got=%d", replicas, updatedWPA.Status.CurrentReplicas)
			}
		})
	}
}
