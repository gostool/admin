package serializer

import (
	"admin/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gmeta"
)

type User struct {
	OrmCommon
	Name     string `json:"name"      ` //
	Password string `json:"password"  ` //
	Nickname string `json:"nickname"  ` //
	RoleId   int    `json:"roleId"    ` //
}

func (u *User) ToData() (data *g.Map) {
	data = &g.Map{
		"id":       u.Id,
		"name":     u.Name,
		"nickname": u.Name,
	}
	return data
}

type UserDetail struct {
	Id       int           `json:"id"`
	Passport string        `json:"passport"`
	RoleId   int           `json:"roleId"`
	RoleMap  map[int]*Role `json:"roleMap"`
}

type UserListInfo struct {
	gmeta.Meta `orm:"table:user"`
	Id         int          `json:"id"`
	Name       string       `json:"name"`
	RoleId     int          `json:"roleId"`
	UserRole   *entity.Role `orm:"with:id=roleId"`
}
