package todo

import (
	"poc-plugin/internal/todo"
	"time"
)

type TodoRequest struct {
	Name string `json:"name"`
	Status string `json:"status"`
}

func (r TodoRequest) ToDomain() todo.Todo {
	return todo.Todo{
		Name: r.Name,
		Status: r.Status,
		CreatedAt: time.Now().String(),
	}
}
