package plugins

var pluginManager Manager

type Plugin struct {
	Name    string
	Enabled bool
	Events  []Event
}

type Manager struct {
	Plugins       []Plugin
	mapPluginsRun map[string]func()
}

var (
	taskEventHandlers = map[string]TaskEventHandler{}
	userEventHandlers = map[string]UserEventHandler{}
)

type TaskEventHandler func(event interface{})
type UserEventHandler func(event interface{})

func RegisterTaskEventHandler(name string, handlerFunc func(event interface{})) {
	taskEventHandlers[name] = handlerFunc
}

func RegisterUserEventHandler(name string, handlerFunc func(event interface{})) {
	userEventHandlers[name] = handlerFunc
}

func HandleUserEvent(userEvent interface{}) {

	for _, handleUserEventFunc := range userEventHandlers {
		handleUserEventFunc(userEvent)
	}
}

func HandleTaskEvent(taskEvent interface{}) {

	for _, handleTaskEventFunction := range taskEventHandlers {
		handleTaskEventFunction(taskEvent)
	}
}
