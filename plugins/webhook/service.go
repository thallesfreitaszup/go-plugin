package webhook

import (
	"encoding/json"
	"log"
	"poc-plugin/internal/configuration/database"
)

type Service struct {
	Repository Repository
}

type Webhook struct {
	Id     int
	URL    string
	Events []string
}

func (w Webhook) ToEntity() database.WebhookDB {
	eventsString, err := json.Marshal(w.Events)
	if err != nil {
		log.Fatal(err)
	}
	return database.WebhookDB{
		URL:    w.URL,
		Events: string(eventsString),
	}
}

func (s Service) Create(webhook Webhook) (int64, error) {
	return s.Repository.Create(webhook.ToEntity())
}

func (s Service) Find() ([]Webhook, error) {
	return s.Repository.Find()
}
