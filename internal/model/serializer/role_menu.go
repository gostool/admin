package serializer

type RoleMenu struct {
	OrmCommon
	RoleId int `json:"roleId"   dc:"角色id" ` // 角色id
	MenuId int `json:"menuId"   dc:"菜单id" ` // 菜单id
}
