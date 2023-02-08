package service

import (
	"bytes"

	"github.com/deyanEnchev/src/model"
)

func GenerateBashScript(tasks []model.Task) string {
	var buffer bytes.Buffer
	buffer.WriteString("#!/usr/bin/env bash\n")
	for _, task := range tasks {
		buffer.WriteString(task.Command + "\n")
	}
	return buffer.String()
}
