package v1

type PageReq struct {
	Page     int `v:"required|min:1#(page)分页号码不能为空|分页号码应当>=1" json:"page" dc:"分页号码" in:"query" d:"1"`
	PageSize int `v:"required|min:10#(pageSize)分页数量不能为空|分页数量应当>=10" json:"pageSize" dc:"分页数量" in:"query" d:"10"`
}
