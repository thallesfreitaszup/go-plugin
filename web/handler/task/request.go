package task

import (
	"poc-plugin/internal/task"
	"time"
)

type TaskRequest struct {
	Name string `json:"name"`
	Status string `json:"status"`
}

func (r TaskRequest) ToDomain() task.Task {
	return task.Task{
		Name: r.Name,
		Status: r.Status,
		CreatedAt: time.Now().String(),
	}
}
