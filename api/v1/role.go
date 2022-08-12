package v1

import (
	"admin/internal/model/serializer"
	"database/sql"
	"github.com/gogf/gf/v2/frame/g"
)

type RoleListReq struct {
	g.Meta `path:"/role/list" method:"get" tags:"RoleService"`
	PageReq
}
type RoleListRes struct {
	Count int                `json:"count"`
	Items []*serializer.Role `json:"items"`
}

type RoleDetailReq struct {
	g.Meta `path:"/role/detail" method:"post" tags:"RoleService"`
	OrmIdReq
}
type RoleDetailRes struct {
	*serializer.Role
}

type RoleUpdateReq struct {
	g.Meta `path:"/role/update" method:"post" tags:"RoleService"`
	OrmIdReq
	RoleAttrReq
}
type RoleUpdateRes struct {
	Id string `json:"name"`
}

type RoleDeleteReq struct {
	g.Meta `path:"/role/delete" method:"post" tags:"RoleService"`
	OrmIdReq
}
type RoleDeleteRes struct {
	Data sql.Result
}

type RoleCreateReq struct {
	g.Meta `path:"/role/create" method:"post" tags:"RoleService"`
	RoleAttrReq
}

type RoleCreateRes struct {
	OrmIdReq
}

type RoleAttrReq struct {
	Name   string `json:"name" dc:"name"`
	Router string `json:"router" dc:"router"`
	Pid    int    `json:"pid" dc:"pid"`
}
