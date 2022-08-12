package v1

import (
	"admin/internal/model/entity"
	"database/sql"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type RoleListReq struct {
	g.Meta `path:"/role/list" method:"get" tags:"RoleService"`
	PageReq
}
type RoleListRes struct {
	Data gdb.Result
}

type RoleDetailReq struct {
	g.Meta `path:"/role/detail" method:"post" tags:"RoleService"`
	OrmIdReq
}
type RoleDetailRes struct {
	Data *entity.Role
}

type RoleUpdateReq struct {
	g.Meta `path:"/role/update" method:"post" tags:"RoleService"`
	OrmIdReq
	Name   string `json:"name" dc:"name"`
	Router string `json:"router" dc:"router"`
	Pid    int    `json:"pid" dc:"pid"`
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
