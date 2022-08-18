// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        int         `json:"id"        ` //
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 修改时间
	DeletedAt *gtime.Time `json:"deletedAt" ` // 删除时间
	IsDeleted int         `json:"isDeleted" ` // 数据的逻辑删除
	Passport  string      `json:"passport"  ` //
	Password  string      `json:"password"  ` //
	Nickname  string      `json:"nickname"  ` //
	RoleId    int         `json:"roleId"    ` //
}
