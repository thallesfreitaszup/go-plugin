package webhook

import (
	"bytes"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"poc-plugin/internal/configuration/database"
)

type Repository struct {
	Orm orm.Ormer
}

func (r Repository) Create(webhook database.WebhookDB) (int64, error) {
	return r.Orm.Insert(&webhook)
}

func (r Repository) Find() ([]Webhook, error) {
	var paramList []database.WebhookDB
	_, err := r.Orm.QueryTable(database.WebhookDB{}).All(&paramList)

	if err != nil {
		return []Webhook{}, err
	}
	return mapToWebHook(paramList), nil
}

func mapToWebHook(webhookDBList []database.WebhookDB) []Webhook {
	var webHookList []Webhook = make([]Webhook, 0)
	for _, webhookDB := range webhookDBList {
		var arrString []string
		json.Unmarshal(bytes.NewBufferString(webhookDB.Events).Bytes(), &arrString)
		webhook := Webhook{
			URL:    webhookDB.URL,
			Events: arrString,
		}
		webHookList = append(webHookList, webhook)
	}
	return webHookList
}
