package todo

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"poc-plugin/internal"
	kafkautils "poc-plugin/internal/kafka"
	"poc-plugin/internal/plugin"
	"poc-plugin/internal/todo"
)
type Handler struct {
	Service todo.Service
	PluginMain plugin.Main
}
func(h Handler) Get(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

type TodoResponse struct {
	Id int64
}

func (h Handler) Post(c echo.Context) error {
	todoRequest := TodoRequest{}
	err := c.Bind(&todoRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	id, err := h.Service.Create(todoRequest.ToDomain())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	todoResponse := TodoResponse {
		Id : id,
	}
	emitEvent(todoRequest.ToDomain(), c.Get(internal.RequestIdValueConstant).(string))
	return c.JSON(http.StatusCreated, todoResponse)
}

func emitEvent(domain todo.Todo, requestIdValue string) {
	kafkautils.Produce(kafkautils.Ctx, todo.TodoCreate, todo.TodoEvent{ Todo: domain, RequestId: requestIdValue} )
}

func (h Handler) Delete(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
func (h Handler) Put(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
