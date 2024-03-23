package src

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"ke-db-migration/config"
)

type Notify struct {
	client *resty.Client
}

func NewNotify() *Notify {
	return &Notify{
		client: resty.New(),
	}
}

func (t *Notify) Qywx(str string) error {
	body := map[string]any{
		"msgtype": "text",
		"text": map[string]any{
			"content": fmt.Sprintf("【%s】%s", config.Config.Notify.Title, str),
		},
	}
	_, err := t.client.R().
		SetHeader("content-type", "application/json").
		SetBody(body).
		Post(config.Config.Notify.Qywx)
	if err != nil {
		return err
	}
	return nil
}
