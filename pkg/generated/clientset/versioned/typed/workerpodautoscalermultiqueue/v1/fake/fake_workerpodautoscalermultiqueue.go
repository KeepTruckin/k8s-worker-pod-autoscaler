/*
Copyright 2019 Practo Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	workerpodautoscalermultiqueuev1 "github.com/practo/k8s-worker-pod-autoscaler/pkg/apis/workerpodautoscalermultiqueue/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeWorkerPodAutoScalerMultiQueues implements WorkerPodAutoScalerMultiQueueInterface
type FakeWorkerPodAutoScalerMultiQueues struct {
	Fake *FakeK8sV1
	ns   string
}

var workerpodautoscalermultiqueuesResource = schema.GroupVersionResource{Group: "k8s.practo.dev", Version: "v1", Resource: "workerpodautoscalermultiqueues"}

var workerpodautoscalermultiqueuesKind = schema.GroupVersionKind{Group: "k8s.practo.dev", Version: "v1", Kind: "WorkerPodAutoScalerMultiQueue"}

// Get takes name of the workerPodAutoScalerMultiQueue, and returns the corresponding workerPodAutoScalerMultiQueue object, and an error if there is any.
func (c *FakeWorkerPodAutoScalerMultiQueues) Get(ctx context.Context, name string, options v1.GetOptions) (result *workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(workerpodautoscalermultiqueuesResource, c.ns, name), &workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueue), err
}

// List takes label and field selectors, and returns the list of WorkerPodAutoScalerMultiQueues that match those selectors.
func (c *FakeWorkerPodAutoScalerMultiQueues) List(ctx context.Context, opts v1.ListOptions) (result *workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueueList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(workerpodautoscalermultiqueuesResource, workerpodautoscalermultiqueuesKind, c.ns, opts), &workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueueList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueueList{ListMeta: obj.(*workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueueList).ListMeta}
	for _, item := range obj.(*workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueueList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested workerPodAutoScalerMultiQueues.
func (c *FakeWorkerPodAutoScalerMultiQueues) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(workerpodautoscalermultiqueuesResource, c.ns, opts))

}

// Create takes the representation of a workerPodAutoScalerMultiQueue and creates it.  Returns the server's representation of the workerPodAutoScalerMultiQueue, and an error, if there is any.
func (c *FakeWorkerPodAutoScalerMultiQueues) Create(ctx context.Context, workerPodAutoScalerMultiQueue *workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueue, opts v1.CreateOptions) (result *workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(workerpodautoscalermultiqueuesResource, c.ns, workerPodAutoScalerMultiQueue), &workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueue), err
}

// Update takes the representation of a workerPodAutoScalerMultiQueue and updates it. Returns the server's representation of the workerPodAutoScalerMultiQueue, and an error, if there is any.
func (c *FakeWorkerPodAutoScalerMultiQueues) Update(ctx context.Context, workerPodAutoScalerMultiQueue *workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueue, opts v1.UpdateOptions) (result *workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(workerpodautoscalermultiqueuesResource, c.ns, workerPodAutoScalerMultiQueue), &workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueue), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWorkerPodAutoScalerMultiQueues) UpdateStatus(ctx context.Context, workerPodAutoScalerMultiQueue *workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueue, opts v1.UpdateOptions) (*workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueue, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(workerpodautoscalermultiqueuesResource, "status", c.ns, workerPodAutoScalerMultiQueue), &workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueue), err
}

// Delete takes name of the workerPodAutoScalerMultiQueue and deletes it. Returns an error if one occurs.
func (c *FakeWorkerPodAutoScalerMultiQueues) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(workerpodautoscalermultiqueuesResource, c.ns, name), &workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueue{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorkerPodAutoScalerMultiQueues) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(workerpodautoscalermultiqueuesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueueList{})
	return err
}

// Patch applies the patch and returns the patched workerPodAutoScalerMultiQueue.
func (c *FakeWorkerPodAutoScalerMultiQueues) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(workerpodautoscalermultiqueuesResource, c.ns, name, pt, data, subresources...), &workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*workerpodautoscalermultiqueuev1.WorkerPodAutoScalerMultiQueue), err
}
