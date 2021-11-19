package authorization

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"poc-plugin/internal"
	"poc-plugin/internal/configuration/database"
	"poc-plugin/plugins"
)

type Handler struct {
	Service Service
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h Handler) Post(c echo.Context) error {
	userRequest := UserRequest{}
	err := c.Bind(&userRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	user, err := h.Service.Create(userRequest.ToEntity())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	requestId := c.Get(internal.RequestIdValueConstant).(string)
	userEvent := createEvent(user, plugins.UserUnauthorized, requestId)
	userResponse := UserResponse{
		Id: user.Id,
	}
	go plugins.HandleUserEvent(userEvent)
	return c.JSON(http.StatusCreated, userResponse)
}

type UserResponse struct {
	Id int
}

func (u UserRequest) ToEntity() database.User {
	return database.User{
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
	}
}
