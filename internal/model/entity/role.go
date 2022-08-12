// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id        int         `json:"id"        ` //
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 修改时间
	DeletedAt *gtime.Time `json:"deletedAt" ` // 删除时间
	IsDeleted int         `json:"isDeleted" ` // 数据的逻辑删除
	Rid       int64       `json:"rid"       ` //
	Name      string      `json:"name"      ` // 角色名称
	Pid       int64       `json:"pid"       ` // 父角色id
	Router    string      `json:"router"    ` // 默认路由
}
