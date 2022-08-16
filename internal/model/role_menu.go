package model

import (
	"admin/internal/consts"
	"github.com/gogf/gf/v2/frame/g"
)

type RoleMenuListInput struct {
	Page     int
	PageSize int
}

type RoleMenuCreateInput struct {
	RoleMenuAttr
}

type RoleMenuDetailInput struct {
	Id int // role_menu id
}

type RoleMenuUpdateInput struct {
	Id int
	RoleMenuAttr
}

type RoleMenuDeleteInput struct {
	Id int
}

func (r *RoleMenuCreateInput) ToMap() (data g.Map) {
	data = g.Map{
		"role_id":    r.RoleId,
		"menu_id":    r.MenuId,
		"is_deleted": consts.CREATED,
	}
	return data
}

func (r *RoleMenuUpdateInput) ToWhereMap() (data g.Map) {
	data = g.Map{
		"id":         r.Id,
		"is_deleted": consts.CREATED,
	}
	return data
}

func (r *RoleMenuUpdateInput) ToMap() (data g.Map) {
	data = g.Map{}
	data["role_id"] = r.RoleId
	data["menu_id"] = r.MenuId
	return data
}

type RoleMenuAttr struct {
	RoleId int
	MenuId int
}
