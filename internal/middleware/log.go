package middleware

import (
	"admin/internal/consts"
	"admin/internal/model"
	"admin/internal/service"
	"bytes"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"io"
	"net/http"
	"time"
)

var SafeHttpMethodSet = gset.NewStrSetFrom([]string{http.MethodGet, http.MethodOptions, http.MethodHead})
var logger = g.Log("debug")

// 日志记录
func Record(r *ghttp.Request) {
	if SafeHttpMethodSet.Contains(r.Method) {
		r.Middleware.Next()
		return
	}
	// 1.记录请求
	now := time.Now()
	record := newOpLogByRequest(r)

	r.Middleware.Next()

	// 2.记录响应: 异步写入
	latency := time.Now().Sub(now)
	record.Latency = int(latency.Microseconds())
	saveOpLog(r, record)
}

func newOpLogByRequest(r *ghttp.Request) (record *model.LogCreateInput) {
	uid := r.GetCtxVar(consts.CtxUserId).Int()
	record = &model.LogCreateInput{}
	record.Ip = r.GetClientIp()
	record.Method = r.Request.Method
	record.Path = r.Request.URL.Path
	record.Agent = r.Request.UserAgent()
	record.UserId = uid
	record.UserName = r.GetCtxVar(consts.CtxUserPassport).String()
	var body []byte
	var err error
	switch record.Method {
	case http.MethodPut, http.MethodPost, http.MethodDelete:
		body, err = io.ReadAll(r.Request.Body)
		ctx := r.GetCtx()
		if err != nil {
			logger.Errorf(ctx, "Record err:%v", err)
		}
		r.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		record.Request = gconv.String(body)
	}
	return record
}

func saveOpLog(r *ghttp.Request, record *model.LogCreateInput) {
	if err := r.GetError(); err != nil {
		record.ErrMsg = err.Error()
	} else {
		record.Response = string(r.Response.Buffer())
	}
	record.Status = r.Response.Status
	ctx := r.GetCtx()
	//service.Log().Create(common, record)
	service.Pool().Invoke(ctx, record)
}
