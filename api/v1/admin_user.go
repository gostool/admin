package v1

import (
	"admin/internal/model/serializer"
	"github.com/gogf/gf/v2/frame/g"
)

type AdminUserDetailReq struct {
	g.Meta `path:"/admin_user/detail" method:"get" tags:"AdminUserService"`
}

type AdminUserDetailRes struct {
	Id       int                      `json:"id"`
	Passport string                   `json:"passport"`
	RoleId   int                      `json:"roleId"`
	RoleMap  map[int]*serializer.Role `json:"roleMap"`
}

type AdminUserListReq struct {
	g.Meta `path:"/admin_user/list" method:"get" tags:"AdminUserService"`
	PageReq
}

type AdminUserListRes struct {
	Id       int                      `json:"id"`
	Passport string                   `json:"passport"`
	RoleId   int                      `json:"roleId"`
	RoleMap  map[int]*serializer.Role `json:"roleMap"`
}

type AdminUserUpdateReq struct {
	g.Meta `path:"/admin_user/update" method:"post" tags:"AdminUserService"`
	OrmIdReq
	AdminUserAttr
}

type AdminUserCreateReq struct {
	g.Meta `path:"/admin_user/create" method:"post" tags:"AdminUserService"`
	AdminUserAttr
}

type AdminUserCreateRes struct {
	Id int `json:"id"`
}

type AdminUserDeleteReq struct {
	g.Meta `path:"/admin_user/delete" method:"post" tags:"AdminUserService"`
	OrmIdReq
}

type AdminUserAttr struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	RoleId   int    `json:"roleId"`
	Nickname string `json:"nickname"`
}
