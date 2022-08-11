package response

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gostool/jsonconv"
)

var logger *glog.Logger

func init() {
	logger = g.Log("debug")
}

// 数据返回通用JSON数据结构
type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码((200:成功, 非200:失败, :错误码))
	Message string      `json:"message"` // 提示信息
	Count   int         `json:"count"`   // 总数量(list接口中返回数据的总量)
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// 标准返回结果数据结构封装。
func Json(ctx context.Context, r *ghttp.Request, code int, message string, count int, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	resp := JsonResponse{
		Code:    code,
		Count:   count,
		Message: message,
		Data:    responseData,
	}
	res, err := json.Marshal(jsonconv.JsonCamelCase{Value: resp})
	if err != nil {
		logger.Error(ctx, err)
		r.Response.WriteJson("")
	} else {
		r.Response.WriteJson(res)
	}
}

func JsonRaw(ctx context.Context, r *ghttp.Request, code int, message string, count int, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	resp := JsonResponse{
		Code:    code,
		Count:   count,
		Message: message,
		Data:    responseData,
	}
	res, err := json.Marshal(resp)
	if err != nil {
		logger.Error(ctx, err)
		r.Response.WriteJson("")
	} else {
		r.Response.WriteJson(res)
	}
}

// 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(ctx context.Context, r *ghttp.Request, err int, msg string, count int, data ...interface{}) {
	switch err {
	case ERR_CODE_BAD:
		logger.Debugf(ctx, "err_code:%v msg:%v req:%v", err, msg, r.GetBodyString())
	case ERR_CODE_ERR:
		logger.Errorf(ctx, "err_code:%v msg:%v data:%v req:%v", err, msg, data, r.GetBodyString())
	}
	Json(ctx, r, err, msg, count, data...)
	r.Exit()
}

func JsonRawExit(ctx context.Context, r *ghttp.Request, err int, msg string, count int, data ...interface{}) {
	JsonRaw(ctx, r, err, msg, count, data...)
	r.Exit()
}

func JsonErrExit(ctx context.Context, r *ghttp.Request, err int, msg string, data ...interface{}) {
	switch err {
	case ERR_CODE_BAD:
		logger.Debugf(ctx, "err_code:%v msg:%v req:%v", err, msg, r.GetBodyString())
	case ERR_CODE_ERR:
		logger.Errorf(ctx, "err_code:%v msg:%v data:%v req:%v", err, msg, data, r.GetBodyString())
	}
	Json(ctx, r, err, msg, 0, data...)
	r.Exit()
}

type Role struct {
	Id     int    `json:"id"`
	Name   string `json:"name""`
	Router string `json:"router"`
}

type User struct {
	Id       int    `json:"id"`
	Passport string `json:"passport"`
	Token    string `json:"token"`
	RoleInfo Role
	RoleId   int `json:"roleId"`
}
