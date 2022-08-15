package v1

type OrmIdReq struct {
	Id int `v:"required|min:1#id不能为空|id应当>=1" json:"id" dc:"id"`
}
type OrmDeleteRes struct {
}

type OrmUpdateRes struct {
}

type OrmSortTypeReq struct {
	// 默认:0 create_at desc 1:crete_at asc 2:update_at desc 3:update asc
	Type int `d:"0" json:"type"`
}
