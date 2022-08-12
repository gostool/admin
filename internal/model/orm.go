package model

import (
	"admin/internal/consts"
	"github.com/gogf/gf/v2/frame/g"
)

type OrmDeleteInput struct {
	Id int // role id
}

func (r *OrmDeleteInput) ToWhereMap() (data g.Map) {
	data = g.Map{
		"id":         r.Id,
		"is_deleted": consts.CREATED,
	}
	return data
}

func (r *OrmDeleteInput) ToMap() (data g.Map) {
	data = g.Map{
		"is_deleted": consts.DELETED,
	}
	return data
}
