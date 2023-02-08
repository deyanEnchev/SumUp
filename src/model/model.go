package model

type Task struct {
	Command  string   `json:"command"`
	Name     string   `json:"name"`
	Requires []string `json:"requires"`
}

type Job struct {
	Tasks []Task `json:"tasks"`
}
