package api

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"poc-plugin/internal/plugin"
	"poc-plugin/internal/todo"
	todoHandler "poc-plugin/web/handler/todo"
)
type Main struct {
	Connection orm.Ormer
}

func (m Main) NewEcho() *echo.Echo{
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	todoHandler := todoHandler.Handler{
		Service: todo.Service{Repository: todo.Repository{ Orm: m.Connection}},
		PluginMain: plugin.Main{
			Echo: e,
		},
	}
	e.GET("/todo", todoHandler.Get)
	e.POST("/todo", todoHandler.Post)
	e.DELETE("/todo", todoHandler.Delete)
	e.PUT("/todo", todoHandler.Put)
	return e
}