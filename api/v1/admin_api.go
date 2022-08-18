package v1

import "github.com/gogf/gf/v2/frame/g"

type AdminApiListReq struct {
	g.Meta `path:"/admin_api/list" method:"get" tags:"AdminApiService"`
	PageReq
}

type AdminApiListRes struct {
	Count int `json:"count" dc:"记录总数"`
	//Items []*serializer.Tpl `json:"items" dc:"条目"`
}

type AdminApiDetailReq struct {
	g.Meta `path:"/admin_api/detail" method:"post" tags:"AdminApiService"`
	OrmIdReq
}

type AdminApiDetailRes struct {
	//*serializer.Tpl
}

type AdminApiUpdateReq struct {
	g.Meta `path:"/admin_api/update" method:"post" tags:"AdminApiService"`
	OrmIdReq
	AdminApiAttrReq
}

type AdminApiUpdateRes struct {
	Id string `json:"id"`
}

type AdminApiDeleteReq struct {
	g.Meta `path:"/admin_api/delete" method:"post" tags:"AdminApiService"`
	OrmIdReq
}

type AdminApiCreateReq struct {
	g.Meta `path:"/admin_api/create" method:"post" tags:"AdminApiService"`
	AdminApiAttrReq
}

type AdminApiCreateRes struct {
	Id string `json:"id"`
}

type AdminApiAttrReq struct {
}
