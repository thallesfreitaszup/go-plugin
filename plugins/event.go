package plugins

import (
	"poc-plugin/internal/configuration/database"
	"poc-plugin/internal/task"
)

type Event string

const (
	TaskCreate       Event = "Task_CREATE"
	TaskUpdate       Event = "Task_UPDATE"
	TaskDelete       Event = "Task_DELETE"
	TaskRead         Event = "Task_READ"
	UserCreate       Event = "USER_CREATE"
	UserDelete       Event = "USER_DELETE"
	UserAuthorized   Event = "USER_AUTHORIZED"
	UserUnauthorized Event = "USER_UNAUTHORIZED"
)

type TaskEvent struct {
	Task      task.Task `json:"task"`
	RequestId string    `json:"requestId"`
	Event     Event     `json:"event"`
}

type UserEvent struct {
	User      database.User `json:"task"`
	RequestId string        `json:"requestId"`
	Event     Event         `json:"event"`
}
