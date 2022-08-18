package serializer

// Role 参考entity
type Role struct {
	OrmCommon
	Name   string `json:"name"     dc:"角色名称" `  // 角色名称
	Pid    int    `json:"pid"      dc:"父角色id" ` // 父角色id
	Router string `json:"router"   dc:"默认路由" `  // 默认路由
}
