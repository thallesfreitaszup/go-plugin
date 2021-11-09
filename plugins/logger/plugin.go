package log

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/labstack/echo/v4"
	"log"
	"poc-plugin/internal/configuration"
	"poc-plugin/internal/configuration/database"
	"poc-plugin/plugins"
	"poc-plugin/plugins/authorization"
	"time"
)
const(
	PluginName = "log"
)

func init() {
	if !isEnabled(){
		return
	}
	manager := configuration.GetDBManager()
	p := Plugin{
		Service: Service{Repository: Repository{Orm: manager}},
		UserService: authorization.Service { Repository: authorization.Repository{Orm: manager} },
	}
	plugins.RegisterUserEventHandler(PluginName, p.handleUserLog)
	plugins.RegisterTaskEventHandler(PluginName, p.handleTaskLog)
	log.Println("Started plugin Logger")
}

func isEnabled() bool {
	return plugins.GetPluginManager().IsPluginEnabled(PluginName)
}


type Plugin struct {
	Echo        *echo.Echo
	Orm         orm.Ormer
	Service     Service
	UserService authorization.Service
}

func (p Plugin) handleTaskLog(taskEventInterface interface{}) {
	taskEvent := (taskEventInterface).(plugins.TaskEvent)
	log.Println("HANDLING TASK LOG", taskEvent.RequestId)
    userLogs, err := p.Service.findByRequestId(taskEvent.RequestId)
	if err != nil {
		log.Println("No previous log on request", err.Error())
	}
	TaskLog := database.TaskLog {
		Action: string(taskEvent.Event),
		RequestId: taskEvent.RequestId,
		Task: &taskEvent.Task,
		User: userLogs.User,
		Timestamp: time.Now(),
	}
	_, err = p.Service.CreateTaskLog(TaskLog)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("FINISHED HANDLING TASK LOG", taskEvent.RequestId)
}

func (p Plugin) handleUserLog(userEventInterface interface {}) {
	userEvent := (userEventInterface).(plugins.UserEvent)

	log.Println("HANDLING USER LOG", userEvent)
	userLog := database.UserLog {
		Action: string(userEvent.Event),
		RequestId: userEvent.RequestId,
		User:      &userEvent.User,
		Timestamp: time.Now(),
	}
	_, err := p.Service.CreateUserLog(userLog)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("FINISHED HANDLING LOG", userEvent)
}
