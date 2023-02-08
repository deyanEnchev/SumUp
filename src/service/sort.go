package service

import (
	"errors"
	"sort"

	"github.com/deyanEnchev/src/model"
)

func SortTasks(job model.Job) ([]model.Task, error) {
	isCycle := cycleDetector(job)
	if isCycle {
		return nil, errors.New("cycle detected")
	}
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

	return job.Tasks, nil
}

func TopologicalSort(j model.Job) ([]model.Task, error) {
	isCycle := cycleDetector(j)
	if isCycle {
		return nil, errors.New("cycle detected")
	}
	taskMap := make(map[string]model.Task)
	for _, task := range j.Tasks {
		taskMap[task.Name] = task
	}

	var sortedTasks []model.Task
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

	return sortedTasks, nil
}

func cycleDetector(job model.Job) bool {
	for i, t1 := range job.Tasks {
		for _, t2 := range job.Tasks[i+1:] {
			isTaskOneRequired := false
			isTaskTwoRequired := false
			for _, req2 := range t2.Requires {
				if t1.Name == req2 {
					isTaskOneRequired = true
				}
			}
			for _, req1 := range t1.Requires {
				if t2.Name == req1 {
					isTaskTwoRequired = true
				}
			}
			if isTaskOneRequired && isTaskTwoRequired {
				return true
			}
		}
	}
	return false
}
