package configuration

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"poc-plugin/internal/configuration/database"
)

var apiManager *echo.Echo
var entityManager orm.Ormer

func init() {
	entityManager = database.CreateDBManager()
	apiManager = ConfigureAPI()

}

func GetAPIManager() *echo.Echo {
	return apiManager
}

func GetDBManager() orm.Ormer {
	return entityManager
}
