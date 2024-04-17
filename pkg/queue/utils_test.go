package queue

import (
	v1 "github.com/practo/k8s-worker-pod-autoscaler/pkg/apis/workerpodautoscalermultiqueue/v1"
	"testing"
)

func TestMinusSet(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		x        map[string]QueueSpec
		y        []v1.Queue
		expected []string
	}{
		{
			name: "missing x element in y",
			x: map[string]QueueSpec{
				"queue1": {
					uri: "queue1",
				},
				"queue2": {
					uri: "queue2",
				},
				"queue3": {
					uri: "queue3",
				},
			},
			y: []v1.Queue{
				{
					URI: "queue1",
				},
				{
					URI: "queue2",
				},
			},
			expected: []string{"queue3"},
		},
		{
			name: "empty set x",
			x:    map[string]QueueSpec{},
			y: []v1.Queue{
				{
					URI: "queue1",
				},
			},
			expected: []string{},
		},
		{
			name: "empty set y",
			x: map[string]QueueSpec{
				"queue1": {
					uri: "queue1",
				},
				"queue2": {
					uri: "queue2",
				},
			},
			y:        []v1.Queue{},
			expected: []string{"queue1", "queue2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinusSet(tt.x, tt.y)
			if len(result) != len(tt.expected) {
				t.Errorf("MinusSet() = %v, expected %v", result, tt.expected)
				return
			}
			for i := 0; i < len(result); i++ {
				if result[i] != tt.expected[i] {
					t.Errorf("MinusSet() = %v, expected %v", result, tt.expected)
				}
			}
		})
	}
}
