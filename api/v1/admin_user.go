package v1

import (
	"admin/internal/model/serializer"
	"github.com/gogf/gf/v2/frame/g"
)

type AdminUserDetailReq struct {
	g.Meta `path:"/admin_user/detail" method:"get" tags:"AdminUserService"`
}

type AdminUserDetailRes struct {
	*serializer.UserDetail
}

type AdminUserListReq struct {
	g.Meta `path:"/admin_user/list" method:"get" tags:"AdminUserService"`
	PageReq
}
type AdminUserListRes struct {
	Count int                    `json:"count" dc:"记录总数"`
	Items []*serializer.UserInfo `json:"items" dc:"条目"`
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
