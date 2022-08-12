// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure of table role for DAO operations like Where/Data.
type Role struct {
	g.Meta    `orm:"table:role, do:true"`
	Id        interface{} //
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
	DeletedAt *gtime.Time // 删除时间
	IsDeleted interface{} // 数据的逻辑删除
	Name      interface{} // 角色名称
	Pid       interface{} // 父角色id
	Router    interface{} // 默认路由
}
