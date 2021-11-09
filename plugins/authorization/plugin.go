package authorization

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"poc-plugin/internal"
	"poc-plugin/internal/kafka"
)


type Plugin struct {
	Echo *echo.Echo
	Orm orm.Ormer
	Handler Handler
}

type UserEvent struct {
	User User `json:"user"`
	RequestId string `json:"requestId"`
}
type Event string
const (
	UserCreate Event = "USER_CREATE"
	UserDelete Event = "USER_DELETE"
	UserAuthorized Event = "USER_AUTHORIZED"
	UserUnauthorized Event = "USER_UNAUTHORIZED"

)


func (p Plugin) Start() {
	p.Handler = Handler{ Service: Service{Repository: Repository{p.Orm}}}
	p.Echo.Use(p.authHandler)
	p.Echo.POST("/user", p.Handler.Post)
	log.Println("Started plugin Authorization")
}


type Request struct {
	Content interface{} `json:"content"`
}


func (p Plugin) authHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := c.Request()
		if request.URL.Path == "/user" {
			return next(c)
		}

		username,password, ok := request.BasicAuth()
		requestId := c.Get(internal.RequestIdValueConstant).(string)
		if !ok {
			kafka.Produce(kafka.Ctx, UserUnauthorized, UserEvent{ User: User { Email: username }, RequestId: requestId})
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		user, err := p.Handler.Service.FindByEmail(username)
		if err != nil {
			log.Println("Error finding user", err.Error())
			kafka.Produce(kafka.Ctx, UserUnauthorized, UserEvent{ User: User { Email: username }, RequestId: requestId})
			return c.JSON(http.StatusInternalServerError, err)

		}
		if user.Password != password {
			log.Println("Password does not match", user)
			kafka.Produce(kafka.Ctx,UserUnauthorized, UserEvent{ User: User { Email: username }, RequestId: requestId})
			return  c.JSON(http.StatusUnauthorized, "Unauthorized")
		}
		kafka.Produce(kafka.Ctx, UserAuthorized, UserEvent{ User: User { Email: username }, RequestId: requestId})
		return next(c)
	}
}