package feishusdk

import (
	"context"
	"testing"
)

func Test_sdkString_sendMsg(t *testing.T) {
	fs := New(
		"https://open.feishu.cn/open-apis/bot/v2/hook/96753e04-d8f7-44d2-94a6-0b2a3a55bb55",
		"hello world!!!",
	)
	err := fs.sendMsg(context.TODO())
	if err != nil {
		t.Log(err)
	}
}
