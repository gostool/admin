package v1

import (
	"admin/internal/model/serializer"
	"github.com/gogf/gf/v2/frame/g"
)

type AdminUserWebReq struct {
	g.Meta `path:"/admin_user/profile" method:"get" tags:"AdminUserService" summary:"get auth user profile"`
}

type AdminUserWebRes struct {
	Id       int                      `json:"id"`
	Passport string                   `json:"passport"`
	RoleId   int                      `json:"roleId"`
	RoleMap  map[int]*serializer.Role `json:"roleMap"`
}
