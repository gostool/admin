package v1

import "github.com/gogf/gf/v2/frame/g"

type RoleListReq struct {
	g.Meta `path:"/role/list" method:"get" tags:"RoleService"`
	PageReq
}
type RoleListRes struct {
}

type RoleDetailReq struct {
	g.Meta `path:"/role/detail" method:"post" tags:"RoleService"`
	Id     int `json:"id"`
}
type RoleDetailRes struct {
}

type RoleUpdateReq struct {
	g.Meta `path:"/role/update" method:"post" tags:"RoleService"`
	Id     int `json:"id"`
}
type RoleUpdateRes struct {
}

type RoleDeleteReq struct {
	g.Meta `path:"/role/delete" method:"post" tags:"RoleService"`
	Id     int `json:"id"`
}
type RoleDeleteRes struct {
}
