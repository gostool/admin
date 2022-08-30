package serializer

import "github.com/gogf/gf/v2/os/gtime"

type RoleMenu struct {
	Id        int         `json:"id"     dc:"id"   `    //
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间" ` // 创建时间
	UpdatedAt *gtime.Time `json:"-" dc:"修改时间" `         // 修改时间
	DeletedAt *gtime.Time `json:"-" dc:"删除时间"`          // 删除时间
	IsDeleted int         `json:"-" `                   // 数据的逻辑删除
	RoleId    int         `json:"roleId"   dc:"角色id" `  // 角色id
	MenuId    int         `json:"menuId"   dc:"菜单id" `  // 菜单id
}
