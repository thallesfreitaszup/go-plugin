package webhook

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	Service Service
}

type WebhookRequest struct {
	URL    string   `json:"url"`
	Events []string `json:"events"`
}

func (w WebhookRequest) Taskmain() Webhook {
	return Webhook{
		URL:    w.URL,
		Events: w.Events,
	}
}

func (h Handler) Post(c echo.Context) error {
	webhookRequest := WebhookRequest{}
	err := c.Bind(&webhookRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	_, err = h.Service.Create(webhookRequest.Taskmain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (h Handler) Find(c echo.Context) error {

	webhookList, err := h.Service.Find()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, webhookList)
}
