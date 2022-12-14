// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleMenu is the golang structure of table role_menu for DAO operations like Where/Data.
type RoleMenu struct {
	g.Meta    `orm:"table:role_menu, do:true"`
	Id        interface{} //
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
	DeletedAt *gtime.Time // 删除时间
	IsDeleted interface{} // 数据的逻辑删除
	RoleId    interface{} // 角色id
	MenuId    interface{} // 菜单id
}
