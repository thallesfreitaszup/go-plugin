package task

import (
	"net/http"
	"poc-plugin/internal"
	"poc-plugin/internal/task"
	"poc-plugin/plugins"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service task.Service
}

func New(manager orm.Ormer, e *echo.Echo) {
	handler := Handler{
		Service: task.Service{Repository: task.RepositoryImpl{Manager: manager}},
	}
	e.GET("/task", handler.Get)
	e.POST("/task", handler.Post)
	e.DELETE("/task/:id", handler.Delete)
	e.PUT("/task/:id", handler.Put)
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
	taskUpdate, err := h.ValidateRequest(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	taskUpdated, err := h.Service.Update(taskUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, taskUpdated.ToResponse())
}

func (h Handler) Get(c echo.Context) error {

	return c.JSON(http.StatusOK, h.Service.Find())
}

func (h Handler) ValidateRequest(c echo.Context) (task.Task, error) {
	taskRequest := TaskRequest{}
	c.Bind(&taskRequest)
	IdString := c.Param("id")
	id, error := strconv.Atoi(IdString)
	if error != nil {
		return task.Task{}, error
	}
	task := taskRequest.ToDomain()
	task.Id = id
	return task, nil
}

type TaskResponse struct {
	Id int
}
