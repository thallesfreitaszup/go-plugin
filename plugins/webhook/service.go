package webhook

import (
	"encoding/json"
	"log"
)
type Service struct {
	Repository Repository
}

type WebhookDB struct {
	Id   int `orm:"auto,column(id)"`
	URL string `orm:"column(url)"`
	Events string `orm:"type(jsonb);column(events)"`
}

type Webhook struct {
Id   int
URL string
Events []string
}

func (w Webhook) ToEntity() WebhookDB {
	eventsString, err  := json.Marshal(w.Events)
	if err != nil {
		log.Fatal(err)
	}
	return WebhookDB{
		URL: w.URL,
		Events: string(eventsString),
	}
}

func (s Service) Create(webhook Webhook) (int64, error) {
	return s.Repository.Create(webhook.ToEntity())
}

func (s Service) Find() ([]Webhook, error) {
	return s.Repository.Find()
}