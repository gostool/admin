package v1

import (
	"admin/internal/model/serializer"
	"github.com/gogf/gf/v2/frame/g"
)

type AdminCasbinListReq struct {
	g.Meta `path:"/admin_casbin/list" method:"get" tags:"AdminCasbinService"`
	RoleId int
}
type AdminCasbinListRes struct {
	Count int               `json:"count" dc:"记录总数"`
	Items []*serializer.Api `json:"items" dc:"条目"`
}

type AdminCasbinUpdateReq struct {
	g.Meta      `path:"/admin_casbin/update" method:"post" tags:"AdminCasbinService"`
	RoleId      int
	ApiInfoList []AdminApiAttr
}
type AdminCasbinUpdateRes struct {
}
