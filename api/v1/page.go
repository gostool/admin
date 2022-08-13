package v1

type PageReq struct {
	Page     int `v:"required" json:"page" dc:"分页号码" in:"query" d:"1"`
	PageSize int `v:"required" json:"pageSize" dc:"分页数量" in:"query" d:"10"`
}
