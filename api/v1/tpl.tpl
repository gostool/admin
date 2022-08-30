type TplListReq struct {
	g.Meta `path:"/tpl/list" method:"get" tags:"TplService"`
	PageReq
	OpTplAttrReqSearch
	OrmSortTypeReq
}
type TplListRes struct {
	Count int               `json:"count" dc:"记录总数"`
	Items []*serializer.Tpl `json:"items" dc:"条目"`
}

type TplDetailReq struct {
	g.Meta `path:"/Tpl/detail" method:"post" tags:"TplService"`
	OrmIdReq
}
type TplDetailRes struct {
	*serializer.Tpl
}

type TplUpdateReq struct {
	g.Meta `path:"/tpl/update" method:"post" tags:"TplService"`
	OrmIdReq
	TplAttrReq
}
type TplUpdateRes struct {
	Id string `json:"id"`
}

type TplDeleteReq struct {
	g.Meta `path:"/tpl/delete" method:"post" tags:"TplService"`
	OrmIdReq
}

type TplCreateReq struct {
	g.Meta `path:"/tpl/create" method:"post" tags:"TplService"`
	TplAttrReq
}

type TplCreateRes struct {
	OrmIdReq
}

type TplTreeReq struct {
	g.Meta `path:"/tpl/tree" method:"get" tags:"TplService"`
	PageReq
}

type TplTreeRes struct {
	Count int                      `json:"count" dc:"记录总数"`
	Items []*serializer.TplDetail `json:"items" dc:"条目"`
}