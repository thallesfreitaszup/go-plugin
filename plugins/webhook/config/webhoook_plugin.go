package config

import (
	"net/rpc"
)
type Task struct {
	Id   int `orm:"auto,column(id)"`
	Name string `orm:"column(name)"`
	CreatedAt string `orm:"column(created_at)"`
	FinishedAt string `orm:"column(finished_at)"`
	Status string `orm:"column(status)"`
}
type Event string
const (
	TaskCreate Event = "Task_CREATE"
	TaskUpdate Event = "Task_UPDATE"
	TaskDelete Event = "Task_DELETE"
	TaskRead  Event = "Task_READ"
)

type TaskEvent struct {
	Task Task `json:"Task"`
	Event Event `json:"event"`
}



type Notifier interface {
	Notify(event TaskEvent) string
}

func (s *NotifierRPCServer) Notify(mapArgs map[string]interface{}, resp *string) error {
	args := mapArgs["data"].(TaskEvent)
	*resp = s.Impl.Notify(args)
	return nil
}

func (g *NotifierRPCClient) Notify(event TaskEvent) string {
	var resp string
	err := g.client.Call("Plugin.Notify", map[string]interface{}{
		"data":   event,
	}, &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}

type NotifierRPCClient struct {
	client *rpc.Client
}

type NotifierRPCServer struct {
	Impl Notifier
}

type NotifierPlugin struct {
	Impl Notifier
}

func (r *NotifierPlugin) Server() (interface{}, error) {
	return &NotifierRPCServer{Impl: r.Impl}, nil
}

func (r *NotifierPlugin) Client(c *rpc.Client) (interface{}, error) {
	return &NotifierRPCClient{client: c}, nil
}