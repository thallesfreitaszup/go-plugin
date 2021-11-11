package log

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	Service Service
}

func (h Handler) GetTaskLogs(c echo.Context) error {
	return 	c.JSON(http.StatusOK, h.Service.GetTasks())
}
