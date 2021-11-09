package plugin

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v2"
	"os"
	"poc-plugin/internal/todo"
)
type Config struct {
	Orm orm.Ormer
	Echo  *echo.Echo
}
type Plugin struct {
	Name string
	Enabled bool
	Events []todo.Event
	Config Config
}

func (p Plugin) Start()  {
	 mapPluginStartFunction[p.Name](p.Config.Echo, p.Config.Orm)
}

type Manager struct {
	Plugins []Plugin
	mapPluginsRun map[string]func()
}
func NewManager() (Manager,error) {
	pluginManager := Manager{}
	log.Info("START:LOAD_PLUGINS")
	data, err := os.ReadFile("plugins.yaml")
	if err != nil {
		return Manager{}, err
	}
	err = yaml.Unmarshal(data, &pluginManager)
	if err != nil {
		return  Manager{}, err
	}
	//pluginManager.RegisterHandlers()
	log.Info("FINISHED:LOAD_PLUGINS", pluginManager.Plugins)
	return pluginManager, nil
}

func (m Manager) handleTodoCreateEvent(todo todo.Todo, event todo.Event) {
	for _, handlerFunction := range todoCreateHandlers {
		handlerFunction(todo, event)
	}
}

func (m Manager) RegisterHandlers() {
	for _, plugin := range m.Plugins {
		 if plugin.Enabled {
			 for _, event := range plugin.Events {
				 m.RegisterEventHandler(event, plugin)
			 }
		 }
	}
}

func (m Manager) RegisterEventHandler(event todo.Event, plugin Plugin) {
	switch event {
	case todo.TodoCreate:
	case todo.TodoUpdate:
	case todo.TodoDelete:
	case todo.TodoRead:
	}
}