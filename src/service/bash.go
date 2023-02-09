package service

import (
	"bytes"

	"github.com/deyanEnchev/src/model"
)

// Generates a bash script from a list of tasks by writing each task's command
// to a buffer and returning the string representation of the buffer.
func GenerateBashScript(tasks []model.Task) string {
	var buffer bytes.Buffer
	buffer.WriteString("#!/usr/bin/env bash\n")
	for _, task := range tasks {
		buffer.WriteString(task.Command + "\n")
	}
	return buffer.String()
}
