package service

import (
	"errors"

	"github.com/deyanEnchev/src/model"
)

func TopologicalSort(j model.Job) ([]model.Task, error) {
	taskMap := make(map[string]model.Task)
	for _, task := range j.Tasks {
		taskMap[task.Name] = task
	}

	var sortedTasks []model.Task
	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	var visit func(taskName string) bool
	visit = func(taskName string) bool {
		if recStack[taskName] {
			return true
		}

		if visited[taskName] {
			return false
		}
		visited[taskName] = true
		recStack[taskName] = true

		task := taskMap[taskName]
		for _, requirement := range task.Requires {
			if visit(requirement) {
				return true
			}
		}

		sortedTasks = append(sortedTasks, task)

		recStack[taskName] = false
		return false
	}

	for _, task := range j.Tasks {
		if visit(task.Name) {
			return nil, errors.New("cycle detected")
		}
	}

	return sortedTasks, nil
}
