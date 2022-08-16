package serializer

import "github.com/gogf/gf/v2/os/gtime"

// Role 参考entity
type Role struct {
	Id        int         `json:"id"     dc:"id"   `    //
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间" ` // 修改时间
	IsDeleted int         `json:"isDeleted" `           // 数据的逻辑删除
	Name      string      `json:"name"     dc:"角色名称" `  // 角色名称
	Pid       int         `json:"pid"      dc:"父角色id" ` // 父角色id
	Router    string      `json:"router"   dc:"默认路由" `  // 默认路由
}
