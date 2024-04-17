package queue

import (
	v1 "github.com/practo/k8s-worker-pod-autoscaler/pkg/apis/workerpodautoscalermultiqueue/v1"
)

// MinusSet returns the elements of x that are not in y
func MinusSet(x map[string]QueueSpec, y []v1.Queue) []string {
	yMap := make(map[string]bool)
	for _, q := range y {
		yMap[q.URI] = true
	}

	result := make([]string, 0)
	for uri, _ := range x {
		if _, ok := yMap[uri]; !ok {
			result = append(result, uri)
		}
	}

	return result
}
