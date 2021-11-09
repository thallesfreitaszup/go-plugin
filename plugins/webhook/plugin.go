package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
	"log"
	"net/http"
	"os"
	kafkautils "poc-plugin/internal/kafka"
)
type Plugin struct {
	Echo *echo.Echo
	Orm orm.Ormer
	Handler Handler
}


func (p Plugin) Start() {

	p.Handler = Handler{ Service: Service{Repository: Repository{p.Orm}}}
	p.Echo.POST("/webhook", p.Handler.Post)
	p.Echo.GET("/webhook", p.Handler.Find)
	log.Println("Started plugin Webhook")
	go p.startListeningEvents()
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
		if err != nil {
			panic("could not read message " + err.Error())

		}
		webhookList, _ := p.Handler.Service.Find()


		for _, webhook := range webhookList {
			p.Notify(webhook, string(msg.Value))
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
	}
}

type Request struct {
	Content interface{} `json:"content"`
}

func (p Plugin) Notify(webhook Webhook, data string ) {
	client := http.Client{}
	requestContent := Request{
		Content: data,
	}
	contentBytes, _ := json.Marshal(requestContent)
	req, _ := http.NewRequest("POST", webhook.URL, bytes.NewReader(contentBytes))
	req.Header.Set("Content-Type","application/json")
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	response.Body.Close()
}

