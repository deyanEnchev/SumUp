package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
)

type Task struct {
	Name     string   `json:"name"`
	Command  string   `json:"command"`
	Requires []string `json:"requires"`
}

type Job struct {
	Tasks []Task `json:"tasks"`
}

func (job *Job) sortTasks() {
	taskIndices := make(map[string]int)
	for i, task := range job.Tasks {
		taskIndices[task.Name] = i
	}
	sort.SliceStable(job.Tasks, func(i, j int) bool {
		a := job.Tasks[i]
		b := job.Tasks[j]
		for _, req := range a.Requires {
			if taskIndices[req] > i {
				return false
			}
		}
		for _, req := range b.Requires {
			if taskIndices[req] > j {
				return true
			}
		}
		return false
	})
}

func (j *Job) TopologicalSort() []Task {
	taskMap := make(map[string]Task)
	for _, task := range j.Tasks {
		taskMap[task.Name] = task
	}

	var sortedTasks []Task
	visited := make(map[string]bool)

	var visit func(taskName string)
	visit = func(taskName string) {
		if visited[taskName] {
			return
		}
		visited[taskName] = true

		task := taskMap[taskName]
		for _, requirement := range task.Requires {
			visit(requirement)
		}

		sortedTasks = append(sortedTasks, task)
	}

	for _, task := range j.Tasks {
		visit(task.Name)
	}

	return sortedTasks
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	var job Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse request body: %v", err), http.StatusBadRequest)
		return
	}

	// sortedTasks := job.TopologicalSort()
	job.sortTasks()
	bashScript := generateBashScript(job.Tasks)

	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte(bashScript))
	// err := json.NewEncoder(w).Encode(script.String())
	if err != nil {
		http.Error(w, "Could not encode response as JSON", http.StatusInternalServerError)
		return
	}
}

func generateBashScript(tasks []Task) string {
	var buffer bytes.Buffer
	buffer.WriteString("#!/usr/bin/env bash\n")
	for _, task := range tasks {
		buffer.WriteString(task.Command + "\n")
	}
	return buffer.String()
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
