package service

import (
	"testing"

	"github.com/deyanEnchev/src/model"
)

func TestGenerateBashScript(t *testing.T) {
	tests := []struct {
		name     string
		tasks    []model.Task
		expected string
	}{
		{
			name: "Single task",
			tasks: []model.Task{
				{Command: "echo 'Hello, World!'"},
			},
			expected: "#!/usr/bin/env bash\necho 'Hello, World!'\n",
		},
		{
			name: "Multiple tasks",
			tasks: []model.Task{
				{Command: "echo 'Hello, World!'"},
				{Command: "ls -l"},
			},
			expected: "#!/usr/bin/env bash\necho 'Hello, World!'\nls -l\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GenerateBashScript(test.tasks)
			if  result != test.expected {
				t.Errorf("Expected %q but got %q", test.expected, result)
			}
		})
	}
}
