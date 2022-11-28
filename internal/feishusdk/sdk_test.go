package feishusdk

import (
	"context"
	"reflect"
	"testing"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

func TestNew(t *testing.T) {
	type args struct {
		webhook string
	}
	tests := []struct {
		name string
		args args
		want *NotifyMessage
	}{
		// TODO: Add test cases.
		{
			name: "01",
			args: args{
				webhook: "https://open.feishu.cn/open-apis/bot/v2/hook/96753e04-d8f7-44d2-94a6-0b2a3a55bb55",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.webhook); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotifyMessage_SendMsg(t *testing.T) {
	resp, _ := gjson.LoadJson(`{"StatusCode":0,"StatusMessage":"success"}"`)
	type fields struct {
		Webhook string
	}
	type args struct {
		ctx context.Context
		msg string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *gjson.Json
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "01",
			fields: fields{
				Webhook: "https://open.feishu.cn/open-apis/bot/v2/hook/96753e04-d8f7-44d2-94a6-0b2a3a55bb55",
			},
			args: args{
				ctx: context.TODO(),
				msg: "hello",
			},
			wantResp: resp,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := New(tt.fields.Webhook)
			gotResp, err := e.SendMsg(tt.args.ctx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Dump(gotResp)
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("SendMsg() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
