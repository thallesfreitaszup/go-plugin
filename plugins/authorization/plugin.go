package authorization

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"poc-plugin/internal"
	"poc-plugin/internal/configuration"
	"poc-plugin/internal/configuration/database"
	"poc-plugin/plugins"
)

const (
	PluginName = "authorization"
)

func init() {
	if !isEnabled() {
		return
	}
	echo := configuration.GetAPIManager()
	manager := configuration.GetDBManager()
	p := Plugin{
		Echo:    echo,
		Handler: Handler{Service: Service{Repository: Repository{Orm: manager}}},
	}
	echo.Use(p.authHandler)
	echo.POST("/user", p.Handler.Post)
	log.Println("Started plugin Authorization")
}

func isEnabled() bool {
	return plugins.GetPluginManager().IsPluginEnabled(PluginName)
}

type Plugin struct {
	Echo    *echo.Echo
	Manager orm.Ormer
	Handler Handler
}

func (p Plugin) authHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := c.Request()
		if request.URL.Path == "/user" {
			return next(c)
		}

		username, password, ok := request.BasicAuth()
		requestId := c.Get(internal.RequestIdValueConstant).(string)
		if !ok {
			user := database.User{Name: username}
			userEvent := createEvent(user, plugins.UserUnauthorized, requestId)
			plugins.HandleUserEvent(userEvent)
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		user, err := p.Handler.Service.FindByEmail(username)
		if err != nil {
			log.Println("Error finding user", err.Error())
			userEvent := createEvent(user, plugins.UserUnauthorized, requestId)
			plugins.HandleUserEvent(userEvent)
			return c.JSON(http.StatusInternalServerError, err)
		}
		if user.Password != password {
			log.Println("Password does not match", user)
			userEvent := createEvent(user, plugins.UserUnauthorized, requestId)
			plugins.HandleUserEvent(userEvent)
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}
		userEvent := createEvent(user, plugins.UserAuthorized, requestId)
		plugins.HandleUserEvent(userEvent)
		return next(c)
	}
}

func createEvent(user database.User, unauthorized plugins.Event, id string) plugins.UserEvent {
	return plugins.UserEvent{
		User:      user,
		Event:     unauthorized,
		RequestId: id,
	}
}
