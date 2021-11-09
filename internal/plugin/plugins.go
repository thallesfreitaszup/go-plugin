package plugin

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/labstack/echo/v4"
	"poc-plugin/internal/todo"
	"poc-plugin/plugins/authorization"
	log "poc-plugin/plugins/logger"
	"poc-plugin/plugins/webhook"
)

var (
	todoReadHandlers              = map[string]TodoReadHandler{}
	todoDeleteHandlers       = map[string]TodoDeleteHandler{}
	todoUpdateHandlers        = map[string]TodoUpdateHandler{}
	todoCreateHandlers          = map[string]TodoCreateHandler{}
)

type TodoReadHandler func(todo todo.Todo, event todo.Event)

type TodoUpdateHandler func(todo todo.Todo, event todo.Event)

type TodoDeleteHandler func(todo todo.Todo, event todo.Event)

type TodoCreateHandler func(todo todo.Todo, event todo.Event)

type StartHandler func(echo * echo.Echo, ormer orm.Ormer)
var mapPluginStartFunction = map[string]StartHandler {
	"webhook": func (echo *echo.Echo, orm orm.Ormer){
		plugin := webhook.Plugin{
			Echo: echo,
			Orm: orm,
		}
		plugin.Start()
	},
	"authorization": func (echo *echo.Echo, orm orm.Ormer){
		plugin := authorization.Plugin{
			Echo: echo,
			Orm: orm,
		}
		plugin.Start()
	},
	"log": func (echo *echo.Echo, orm orm.Ormer){
		plugin := log.Plugin{
			Echo: echo,
			Orm: orm,
		}
		plugin.Start()
	},
}

type Notifier interface {
	Notify(event todo.TodoEvent)
}
