package service

import (
	"errors"
	"testing"

	"github.com/deyanEnchev/src/model"
)

func TestTopologicalSort(t *testing.T) {
	tests := []struct {
		name     string
		job      model.Job
		expected []model.Task
		err      error
	}{
		{
			name: "No dependencies",
			job: model.Job{
				Tasks: []model.Task{
					{Name: "A", Requires: []string{}},
					{Name: "B", Requires: []string{}},
					{Name: "C", Requires: []string{}},
				},
			},
			expected: []model.Task{
				{Name: "A", Requires: []string{}},
				{Name: "B", Requires: []string{}},
				{Name: "C", Requires: []string{}},
			},
			err: nil,
		},
		{
			name: "Dependencies 1",
			job: model.Job{
				Tasks: []model.Task{
					{Name: "A", Requires: []string{"C"}},
					{Name: "B", Requires: []string{"A"}},
					{Name: "C", Requires: []string{}},
				},
			},
			expected: []model.Task{
				{Name: "C", Requires: []string{}},
				{Name: "A", Requires: []string{"C"}},
				{Name: "B", Requires: []string{"A"}},
			},
			err: nil,
		},
		{
			name: "Dependencies 2",
			job: model.Job{
				Tasks: []model.Task{
					{Name: "A", Requires: []string{"B"}},
					{Name: "B", Requires: []string{"C"}},
					{Name: "C", Requires: []string{}},
				},
			},
			expected: []model.Task{
				{Name: "C", Requires: []string{}},
				{Name: "B", Requires: []string{"C"}},
				{Name: "A", Requires: []string{"B"}},
			},
			err: nil,
		},
		{
			name: "Dependencies 3",
			job: model.Job{
				Tasks: []model.Task{
					{Name: "A", Requires: []string{"B", "C"}},
					{Name: "B", Requires: []string{}},
					{Name: "C", Requires: []string{}},
				},
			},
			expected: []model.Task{
				{Name: "B", Requires: []string{}},
				{Name: "C", Requires: []string{}},
				{Name: "A", Requires: []string{"B", "C"}},
			},
			err: nil,
		},
		{
			name: "Cycle 1",
			job: model.Job{
				Tasks: []model.Task{
					{Name: "A", Requires: []string{"B"}},
					{Name: "B", Requires: []string{"A"}},
				},
			},
			expected: nil,
			err:      errors.New("cycle detected"),
		},
		{
			name: "Cycle 2",
			job: model.Job{
				Tasks: []model.Task{
					{Name: "A", Requires: []string{"B"}},
					{Name: "B", Requires: []string{"C"}},
					{Name: "C", Requires: []string{"A"}},
				},
			},
			expected: nil,
			err:      errors.New("cycle detected"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := TopologicalSort(test.job)
			if err != nil {
				if err.Error() != test.err.Error() {
					t.Errorf("Expected error %q but got %q", test.err, err)
				}
				return
			}
			if len(result) != len(test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, result)
				return
			}
			for i, task := range result {
				if task.Name != test.expected[i].Name {
					t.Errorf("Expected task %v but got %v", test.expected[i], task)
					return
				}
			}
		})
	}
}
