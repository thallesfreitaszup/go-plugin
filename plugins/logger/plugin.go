package log

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	kafkautils "poc-plugin/internal/kafka"
	"poc-plugin/internal/todo"
	"poc-plugin/plugins/authorization"
	"time"
)
type Plugin struct {
	Echo *echo.Echo
	Orm orm.Ormer
	Service Service
	UserService authorization.Service
}
type Event string

const (
	UserCreate Event = "USER_CREATE"
	UserAuthorized Event = "USER_AUTHORIZED"
	UserUnauthorized Event = "USER_UNAUTHORIZED"
	TodoCreate Event = "TODO_CREATE"
	TodoUpdate Event = "TODO_UPDATE"
	TodoDelete Event = "TODO_DELETE"
	TodoRead  Event = "TODO_READ"
)

func (p Plugin) Start() {

	p.Service =  Service { Repository: Repository{p.Orm}}
	p.UserService =  authorization.Service { Repository: authorization.Repository{Orm: p.Orm}}
	p.Echo.Use(requestIdMiddleWare)
	log.Println("Started plugin Log ")
	go p.startListeningEvents()
}

func requestIdMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		id := uuid.New()

		c.Set("RequestIdValueConstant", id)

		return next(c)
	}
}

func(p Plugin) startListeningEvents() {

	// create a new logger that outputs to stdout
	// and has the `kafka reader` prefix
	l := log.New(os.Stdout, "kafka reader: ", 0)
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkautils.BrokerAddress},
		Topic:   kafkautils.Topic,
		GroupID: "my-group",
		// assign the logger to the reader
		Logger: l,
	})
	for {

		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(kafkautils.Ctx)
		fmt.Println("Log Plugin: received: ", string(msg.Value))
		if err != nil {
			panic("could not read message " + err.Error())

		}
		p.handleLogEvent(msg.Key, msg.Value)

		// after receiving the message, log its value
	}
}

func(p Plugin) handleLogEvent(key []byte, value []byte) {
	var event Event
	var todoEvent todo.TodoEvent
	var userEvent authorization.UserEvent
	err := json.Unmarshal(key, &event)
	if err != nil {
		log.Fatal("Error reading event")
	}
	switch event {
	case TodoCreate, TodoUpdate, TodoDelete, TodoRead:
		json.Unmarshal(value, &todoEvent)
		p.handleTodoLog(event, todoEvent)
	case UserAuthorized, UserUnauthorized, UserCreate:
		json.Unmarshal(value, &userEvent)
		p.handleUserLog(event, userEvent)
	default:
		log.Fatalln("Unsupported event")
	}
}

type TodoLog struct {
	Id int `orm:"auto;column(id)"`
	Action Event `orm:"column(action)"`
	RequestId string `orm:"column(request_id)"`
	User *authorization.User `orm:"null;rel(fk);on_delete(set_null)"`
	Todo *todo.Todo `orm:"null;rel(fk);on_delete(set_null)"`
	Timestamp  time.Time `orm:"column(timestamp)"`
}

type UserLog struct {
	Id int `orm:"auto;column(id)"`
	Action Event `orm:"column(action)"`
	RequestId string `orm:"column(request_id)"`
	User *authorization.User `orm:"null;rel(fk);on_delete(set_null)"`
	Timestamp  time.Time `orm:"column(timestamp)"`
}

func (p Plugin) handleTodoLog(action Event, event todo.TodoEvent) {
	var user  authorization.User
	log.Println("HANDLING TODO LOG", event)
    todoLogs, err := p.Service.findByRequestId(event.RequestId)
	for _, todoLog := range todoLogs {
		if todoLog.User.Id != 0 {
			user, err = p.UserService.FindById(todoLog.User.Id)
			if err == nil {
				break
			}else {
				log.Fatal("Error finding user log", err.Error())
			}
		}
	}
	if err != nil {
		log.Println("No previous log on request")
	}
	todoLog := TodoLog {
		Action: action,
		RequestId: event.RequestId,
		Todo: &event.Todo,
		User: &user,
		Timestamp: time.Now(),
	}
	p.Service.CreateTodoLog(todoLog)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (p Plugin) handleUserLog(key Event, event authorization.UserEvent) {
	log.Println("HANDLING USER LOG", event)
	userLog := UserLog {
		Action: key,
		RequestId: event.RequestId,
		User: &event.User,
		Timestamp: time.Now(),
	}
	_, err := p.Service.CreateUserLog(userLog)
	if err != nil {
		log.Fatal(err.Error())
	}
}
