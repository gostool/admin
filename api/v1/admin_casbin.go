package v1

import (
	"admin/internal/model/serializer"
	"github.com/gogf/gf/v2/frame/g"
)

type AdminCasbinListReq struct {
	g.Meta `path:"/admin_casbin/list" method:"get" tags:"AdminCasbinService"`
	RoleId int `v:"required|min:1#id不能为空|id应当>=1" json:"roleId"`
}
type AdminCasbinListRes struct {
	Count int                     `json:"count" dc:"记录总数"`
	Items []*serializer.CasbinApi `json:"items" dc:"条目"`
}

type AdminCasbinUpdateReq struct {
	g.Meta      `path:"/admin_casbin/update" method:"post" tags:"AdminCasbinService"`
	RoleId      int               `v:"required|min:1#id不能为空|id应当>=1" json:"roleId"`
	ApiInfoList []AdminCasbinAttr `json:"apiInfoList"`
}
type AdminCasbinUpdateRes struct {
}

type AdminCasbinAttr struct {
	Path   string `v:"required|length:1,100#请输入path|path长度应当在:1到:100之间" json:"path"   dc:"请求路径"`   // 请求路径
	Method string `v:"required|length:1,10#请输入method|method长度应当在:1到:10之间" json:"method" dc:"请求方式"` // 请求方式
}
