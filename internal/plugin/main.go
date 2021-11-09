package plugin

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/labstack/echo/v4"
	todo "poc-plugin/internal/todo"
)

type Main struct {
	Echo *echo.Echo
	Manager Manager
	Orn orm.Ormer
}

func (m Main) HandleEvent(data interface{}, event todo.Event) {
	switch event {
	case todo.TodoCreate:
		todo := data.(todo.Todo)
		m.Manager.handleTodoCreateEvent(todo, event)
  }

}

func (m Main) Start() {
	for _, plugin := range m.Manager.Plugins {
		if plugin.Enabled {
			plugin.Config = Config{
				Echo: m.Echo,
				Orm: m.Orn,
			}
			plugin.Start()
		}
	}
}
