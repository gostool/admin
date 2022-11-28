package feishusdk

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type NotifyMessage struct {
	Webhook string
}

func New(webhook string) *NotifyMessage {
	sdk := &NotifyMessage{
		Webhook: webhook,
	}
	return sdk
}

func (e *NotifyMessage) SendMsg(ctx context.Context, msg string) (resp *gjson.Json, err error) {
	data := g.Map{
		"msg_type": "text",
		"content": g.Map{
			"text": msg,
		},
	}
	r, err := g.Client().Post(
		ctx,
		e.Webhook,
		data,
	)
	if err != nil {
		return resp, err
	}
	defer r.Close()
	resp, err = gjson.LoadJson(r.ReadAllString())
	if err != nil {
		return resp, err
	}
	return resp, nil
}
