package v1

type OrmIdReq struct {
	Id int `v:"required|min:1#id不能为空|id应当>=1" json:"id" dc:"id"`
}
type OrmDeleteRes struct {
}

type OrmUpdateRes struct {
}
