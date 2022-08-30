package model

import (
	"admin/internal/consts"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type RoleListInput struct {
	Page     int
	PageSize int
}

type RoleCreateInput struct {
	RoleAttr
}

type RoleDetailInput struct {
	Id int // role id
}

type RoleUpdateInput struct {
	Id int
	RoleAttr
}

type RoleDeleteInput struct {
	Id int
}

type RoleAttr struct {
	// 角色名
	Name string `json:"name"`
	// 角色默认前端路由
	Router string `json:"router"`
	// 父角色id
	Pid int `json:"pid"`
}

func (r *RoleCreateInput) ToMap() (data g.Map) {
	if r.Router == "" {
		r.Router = "dashboard"
	}
	data = g.Map{
		"name":       r.Name,
		"router":     r.Router,
		"is_deleted": consts.CREATED,
	}
	if r.Pid != -1 {
		data["pid"] = r.Pid
	}
	return data
}

func (r *RoleUpdateInput) ToWhereMap() (data g.Map) {
	data = g.Map{
		"id":         r.Id,
		"is_deleted": consts.CREATED,
	}
	return data
}

func (r *RoleUpdateInput) ToMap() (data g.Map) {
	data = g.Map{}
	if r.Pid != 0 {
		data["pid"] = r.Pid
	}
	if r.Pid == -1 {
		data["pid"] = gdb.Raw("NULL")
	}
	if r.Name != "" {
		data["name"] = r.Name
	}
	if r.Router != "" {
		data["router"] = r.Router
	}
	return data
}
