package main

import (
	"poc-plugin/internal/configuration"
	_ "poc-plugin/plugins/plugin-imports"
	taskHandler "poc-plugin/web/handler/task"
)

func main() {
	apiManager := configuration.GetAPIManager()
	dbManager := configuration.GetDBManager()
	taskHandler.New(dbManager, apiManager)
	apiManager.Logger.Fatal(apiManager.Start(":8080"))
}
