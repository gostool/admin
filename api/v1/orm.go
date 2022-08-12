package v1

type OrmIdReq struct {
	Id int `v:"required" json:"id" dc:"id"`
}
type OrmDeleteRes struct {
}

type OrmUpdateRes struct {
}
