package serializer

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

type User struct {
	Id        int         `json:"id"     dc:"id"   `    //
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间" ` // 创建时间
	UpdatedAt *gtime.Time `json:"-" dc:"修改时间" `         // 修改时间
	DeletedAt *gtime.Time `json:"-" dc:"删除时间"`          // 删除时间
	IsDeleted int         `json:"-" `                   // 数据的逻辑删除
	Name      string      `json:"name"      `           //
	Password  string      `json:"password"  `           //
	Nickname  string      `json:"nickname"  `           //
	RoleId    int         `json:"roleId"    `           //
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
	Name     string        `json:"name"`
	Nickname string        `json:"nickname"`
	RoleId   int           `json:"roleId"`
	RoleMap  map[int]*Role `json:"roleMap"`
}

type UserInfo struct {
	gmeta.Meta `orm:"table:user"`
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Nickname   string    `json:"nickname"`
	RoleId     int       `json:"roleId"`
	RoleMap    *RoleWith `orm:"with:id=roleId"`
}
