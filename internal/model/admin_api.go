package model

import (
	"admin/internal/consts"
	"github.com/gogf/gf/v2/frame/g"
)

type AdminApiListInput struct {
	Page     int
	PageSize int
}

type AdminApiCreateInput struct {
	AdminApiAttr
}

type AdminApiDetailInput struct {
	Id int // id
}

type AdminApiUpdateInput struct {
	Id int
	AdminApiAttr
}

type AdminApiDeleteInput struct {
	Id int
}

type AdminApiAttr struct {
	Group  string `json:"group"  dc:"分组"`   // 分组
	Path   string `json:"path"   dc:"请求路径"` // 请求路径
	Detail string `json:"detail" dc:"简介"`   // 简介
	Method string `json:"method" dc:"请求方式"` // 请求方式
}

func (r *AdminApiCreateInput) ToMap() (data g.Map) {
	data = g.Map{
		"group":      r.Group,
		"path":       r.Path,
		"detail":     r.Detail,
		"method":     r.Method,
		"is_deleted": consts.CREATED,
	}
	return data
}

func (r *AdminApiUpdateInput) ToWhereMap() (data g.Map) {
	data = g.Map{
		"id":         r.Id,
		"is_deleted": consts.CREATED,
	}
	return data
}

func (r *AdminApiUpdateInput) ToMap() (data g.Map) {
	data = g.Map{}
	data["group"] = r.Group
	data["path"] = r.Path
	data["detail"] = r.Detail
	data["method"] = r.Method
	return data
}
