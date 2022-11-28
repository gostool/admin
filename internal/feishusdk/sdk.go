package feishusdk

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
)

func New(url, msg string) *sdkString {
	sdk := &sdkString{
		url: url,
		data: g.Map{
			"msg_type": "text",
			"content": g.Map{
				"text": msg,
			},
		},
	}
	return sdk
}

type sdkString struct {
	url  string
	data g.Map
}

func (e *sdkString) sendMsg(ctx context.Context) error {
	r, err := g.Client().Post(
		ctx,
		e.url,
		e.data,
	)
	if err != nil {
		return err
	}
	defer r.Close()
	fmt.Println(r.ReadAllString())
	return nil
}
