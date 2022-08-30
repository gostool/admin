package model

import (
	"admin/internal/consts"
	"github.com/gogf/gf/v2/frame/g"
)

type OpTplAttrSearch struct {
	// 响应状态: 200/500/400
	Status int `json:"status"`
	// 请求方法: get/post
	Method string `json:"method"`
	// 请求路径
	Path string `json:"path"`
}

type OrmSortType struct {
	// 排序字段: "create_at/update_at desc"
	//Filed string `json:"filed"`
	Type int `json:"type"`
}

type TplListInput struct {
	Page     int
	PageSize int
}

type TplSearchInput struct {
	Page     int
	PageSize int
	OpTplAttrSearch
	OrmSortType
}

type TplCreateInput struct {
	TplAttr
}

type TplDetailInput struct {
	Id int // log id
}

type TplUpdateInput struct {
	Id int
	LogAttr
}

type TplDeleteInput struct {
	Id int
}

type TplAttr struct {
	UserId   int    `json:"userId"    ` //
	Status   int    `json:"status"    ` // 状态
	Extra    string `json:"extra"     ` //
	Ip       string `json:"ip"        ` // 请求ip
	Path     string `json:"path"      ` // 请求路径
	Method   string `json:"method"    ` // 请求方法
	Request  string `json:"request"   ` // 请求参数
	Response string `json:"response"  ` // 响应内容
	Agent    string `json:"agent"     ` // 代理
	Latency  int    `json:"latency"   ` // 延迟
	ErrMsg   string `json:"errMsg"    ` // 错误信息
	UserName string `json:"userName"  ` // 用户名
}

func (r *TplCreateInput) ToMap() (data g.Map) {
	data = g.Map{
		"userId":     r.UserId,
		"status":     r.Status,
		"extra":      r.Extra,
		"ip":         r.Ip,
		"path":       r.Path,
		"method":     r.Method,
		"request":    r.Request,
		"response":   r.Response,
		"agent":      r.Agent,
		"errMsg":     r.ErrMsg,
		"userName":   r.UserName,
		"is_deleted": consts.CREATED,
	}
	return data
}

func (r *TplUpdateInput) ToWhereMap() (data g.Map) {
	data = g.Map{
		"id":         r.Id,
		"is_deleted": consts.CREATED,
	}
	return data
}

func (r *TplUpdateInput) ToMap() (data g.Map) {
	data = g.Map{}
	data["userId"] = r.UserId
	data["status"] = r.Status
	data["extra"] = r.Extra
	data["ip"] = r.Ip
	data["path"] = r.Path
	data["method"] = r.Method
	data["request"] = r.Request
	data["response"] = r.Response
	data["agent"] = r.Agent
	data["errMsg"] = r.ErrMsg
	data["userName"] = r.UserName
	return data
}
