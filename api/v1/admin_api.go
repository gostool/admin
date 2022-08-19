package v1

import (
	"admin/internal/model/serializer"
	"github.com/gogf/gf/v2/frame/g"
)

type AdminApiListReq struct {
	g.Meta `path:"/admin_api/list" method:"get" tags:"AdminApiService"`
	PageReq
}

type AdminApiListRes struct {
	Count int               `json:"count" dc:"记录总数"`
	Items []*serializer.Api `json:"items" dc:"条目"`
}

type AdminApiDetailReq struct {
	g.Meta `path:"/admin_api/detail" method:"post" tags:"AdminApiService"`
	OrmIdReq
}

type AdminApiDetailRes struct {
	*serializer.Api
}

type AdminApiUpdateReq struct {
	g.Meta `path:"/admin_api/update" method:"post" tags:"AdminApiService"`
	OrmIdReq
	AdminApiAttr
}

type AdminApiDeleteReq struct {
	g.Meta `path:"/admin_api/delete" method:"post" tags:"AdminApiService"`
	OrmIdReq
}

type AdminApiCreateReq struct {
	g.Meta `path:"/admin_api/create" method:"post" tags:"AdminApiService"`
	AdminApiAttr
}

type AdminApiAttr struct {
	Group  string `v:"required|length:1,100#请输入group|group长度应当在:1到:100之间" json:"group"  dc:"分组"`     // 分组
	Path   string `v:"required|length:1,100#请输入path|path长度应当在:1到:100之间" json:"path"   dc:"请求路径"`     // 请求路径
	Detail string `v:"required|length:1,1000#请输入detail|detail长度应当在:1到:1000之间" json:"detail" dc:"简介"` // 简介
	Method string `v:"required|length:1,10#请输入method|method长度应当在:1到:10之间" json:"method" dc:"请求方式"`   // 请求方式
}
