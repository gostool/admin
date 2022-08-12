package v1

type PageReq struct {
	Page     int `v:"required" json:"page" dc:"page" in:"query" d:"1"`
	PageSize int `v:"required" json:"pageSize" dc:"page size" in:"query" d:"10"`
}
