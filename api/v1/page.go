package v1

type PageReq struct {
	page     int `v:"required" json:"page" dc:"page" in:"query" d:"1"`
	pageSize int `v:"required" json:"pageSize" dc:"page size" in:"query" d:"10"`
}
