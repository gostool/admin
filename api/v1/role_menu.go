package v1

import (
	"admin/internal/model/serializer"
	"database/sql"
	"github.com/gogf/gf/v2/frame/g"
)

type RoleMenuListReq struct {
	g.Meta `path:"/role_menu/list" method:"get" tags:"RoleMenuService"`
	PageReq
}
type RoleMenuListRes struct {
	Count int                    `json:"count" dc:"记录总数"`
	Items []*serializer.RoleMenu `json:"items" dc:"条目"`
}

type RoleMenuDetailReq struct {
	g.Meta `path:"/role_menu/detail" method:"post" tags:"RoleMenuService"`
	OrmIdReq
}
type RoleMenuDetailRes struct {
	*serializer.RoleMenu
}

type RoleMenuUpdateReq struct {
	g.Meta `path:"/role_menu/update" method:"post" tags:"RoleMenuService"`
	OrmIdReq
	RoleMenuAttr
}
type RoleMenuUpdateRes struct {
	Id string `json:"id"`
}

type RoleMenuDeleteReq struct {
	g.Meta `path:"/role_menu/delete" method:"post" tags:"RoleMenuService"`
	OrmIdReq
}
type RoleMenuDeleteRes struct {
	Data sql.Result
}

type RoleMenuCreateReq struct {
	g.Meta `path:"/role_menu/create" method:"post" tags:"RoleMenuService"`
	RoleMenuAttr
}

type RoleMenuCreateRes struct {
	OrmIdReq
}

type RoleMenuAttr struct {
	RoleId int
	MenuId int
}
