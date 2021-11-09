package main

import (
	"log"
	"poc-plugin/internal/configuration"
	"poc-plugin/internal/plugin"
	"poc-plugin/web/api"
)

func main() {
	var connection = configuration.GetDBConnection()
	pluginManager, err := plugin.NewManager()
	if err != nil {
		log.Fatalln(err)
	}
	apiMain := api.Main{
		Connection: connection,
	}
	e:= apiMain.NewEcho()
	pluginMain := plugin.Main{
		Echo: e,
		Manager: pluginManager,
		Orn: connection,
	}
	pluginMain.Start()
	e.Logger.Fatal(e.Start(":8080"))
}

//// Handler
//func hello(c echo.Context) error {
//	return c.String(http.StatusOK, "Hello, World!")
//}