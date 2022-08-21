package v1

type OrmIdReq struct {
	Id int `v:"required|min:1#id不能为空|id应当>=1" json:"id" dc:"id"`
}

type OrmSortTypeReq struct {
	// 默认:0 create_at desc 1:crete_at asc 2:update_at desc 3:update asc
	Type int `d:"0" json:"type"`
}

//----------上面req  下面res

type OrmDeleteRes struct {
}

type OrmUpdateRes struct {
}

type OrmIdRes struct {
	Id int `json:"id" dc:"id"`
}

type OrmOrderParamsReq struct {
	OrderKey string `json:"orderKey" d:"id"`
	Desc     int    `json:"desc" d:"0"`
}
