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

type AdminUserRegisterReq struct {
	g.Meta `path:"/admin_user/register" method:"post" tags:"AdminUserService"`
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
	Name     string `v:"required|length:5,16#请输入用户名|用户名称长度应当在:5到:16之间" dc:"Your name" json:"name"`
	Password string `v:"required|length:5,16#请输入确认密码|密码长度应当在:5到:16之间" dc:"Your password" json:"password"`
	Nickname string `v:"required|length:5,16#请输入昵称|昵称长度应当在:5到:16之间" dc:"Your nickname" json:"nickname"`
	RoleId   int    `v:"required|min:1#(roleId)角色Id不能为空|角色Id应当>=1" dc:"Your roleId" json:"roleId"`
}
