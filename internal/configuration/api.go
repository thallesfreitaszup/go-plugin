package configuration

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigureAPI() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(requestIdMiddleWare)
	return e
}


func requestIdMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		id := uuid.New()

		c.Set("RequestIdValueConstant", id.String())

		return next(c)
	}
}
