package serializer

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role 参考entity
type Role struct {
	Id        int         `json:"id"     dc:"id"   `    //
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间" ` // 创建时间
	UpdatedAt *gtime.Time `json:"-" dc:"修改时间" `         // 修改时间
	DeletedAt *gtime.Time `json:"-" dc:"删除时间"`          // 删除时间
	IsDeleted int         `json:"-" `                   // 数据的逻辑删除
	Name      string      `json:"name"     dc:"角色名称" `  // 角色名称
	Pid       int         `json:"pid"      dc:"父角色id" ` // 父角色id
	Router    string      `json:"router"   dc:"默认路由" `  // 默认路由
}

type RoleWith struct {
	g.Meta    `orm:"table:role, do:true"`
	Id        int         `json:"id"     dc:"id"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间" ` // 创建时间
	UpdatedAt *gtime.Time `json:"-" dc:"修改时间" `         // 修改时间
	DeletedAt *gtime.Time `json:"-" dc:"删除时间"`          // 删除时间
	IsDeleted int         `json:"-" `                   // 数据的逻辑删除
	Name      string      `json:"name"     dc:"角色名称" `  // 角色名称
	Pid       int         `json:"pid"      dc:"父角色id" ` // 父角色id
	Router    string      `json:"router"   dc:"默认路由" `  // 默认路由
}

type RoleDetail struct {
	Role
	Children []*RoleDetail `json:"children"`
}
