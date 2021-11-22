package webhook

import (
	"bytes"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/labstack/echo/v4"
	logstack "github.com/labstack/gommon/log"
	"log"
	"net/http"
	"poc-plugin/internal/configuration"
	"poc-plugin/plugins"
)

const (
	pluginName = "webhook"
)

var main Main

func init() {
	if !isEnabled() {
		return
	}
	entityManager := configuration.GetDBManager()
	handler := Handler{Service: Service{Repository: Repository{entityManager}}}
	main = Main{
		ApiManager:    configuration.GetAPIManager(),
		EntityManager: entityManager,
		Handler:       handler,
	}
	main.ApiManager.POST("/webhook", main.Handler.Post)
	main.ApiManager.GET("/webhook", main.Handler.Find)
	plugins.RegisterTaskEventHandler(pluginName, main.Notify)
	plugins.RegisterUserEventHandler(pluginName, main.Notify)
	logstack.Info("Started Plugin Webhook")
}

func isEnabled() bool {
	return plugins.GetPluginManager().IsPluginEnabled(pluginName)
}

type Main struct {
	ApiManager    *echo.Echo
	EntityManager orm.Ormer
	Handler       Handler
}

type Request struct {
	Content interface{} `json:"content"`
}

func (m Main) Notify(event interface{}) {

	data, _ := json.Marshal(event)

	webhookList, err := m.Handler.Service.Find()
	if err != nil {
		logstack.Error("Failed to find webhooks")
	}
	for _, webhook := range webhookList {
		m.sendRequest(string(data), webhook.URL)
	}
}

func (m Main) sendRequest(data string, url string) {
	requestContent := Request{
		Content: data,
	}
	client := http.Client{}
	contentBytes, _ := json.Marshal(requestContent)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(contentBytes))
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	response.Body.Close()
}
