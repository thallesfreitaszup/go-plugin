package task

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/labstack/echo/v4"
	"net/http"
	"poc-plugin/internal"
	"poc-plugin/internal/task"
	"poc-plugin/plugins"
)

type Handler struct {
	Service task.Service
}

func (h Handler) Get(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

type TaskResponse struct {
	Id int
}

func New(manager orm.Ormer, e *echo.Echo) {
	handler := Handler{
		Service: task.Service{Repository: task.RepositoryImpl{Manager: manager}},
	}
	e.GET("/task", handler.Get)
	e.POST("/task", handler.Post)
	e.DELETE("/task", handler.Delete)
	e.PUT("/task", handler.Put)
}

func (h Handler) Post(c echo.Context) error {
	TaskRequest := TaskRequest{}
	err := c.Bind(&TaskRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	task, err := h.Service.Create(TaskRequest.ToDomain())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	TaskResponse := TaskResponse{
		Id: task.Id,
	}
	handleEvent(plugins.TaskCreate, task, c.Get(internal.RequestIdValueConstant).(string))
	return c.JSON(http.StatusCreated, TaskResponse)
}

func handleEvent(event plugins.Event, domain task.Task, requestIdValue string) {
	taskEvent := plugins.TaskEvent{
		Task:      domain,
		Event:     event,
		RequestId: requestIdValue,
	}
	plugins.HandleTaskEvent(taskEvent)
}

func (h Handler) Delete(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
func (h Handler) Put(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
