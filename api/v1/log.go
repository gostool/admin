package v1

import (
	"admin/internal/model/serializer"
	"database/sql"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type LogListReq struct {
	g.Meta `path:"/log/list" method:"get" tags:"LogService"`
	PageReq
}
type LogListRes struct {
	Count int               `json:"count" dc:"记录总数"`
	Items []*serializer.Log `json:"items" dc:"条目"`
}

type LogDetailReq struct {
	g.Meta `path:"/log/detail" method:"post" tags:"LogService"`
	OrmIdReq
}
type LogDetailRes struct {
	*serializer.Log
}

type LogUpdateReq struct {
	g.Meta `path:"/log/update" method:"post" tags:"LogService"`
	OrmIdReq
	LogAttrReq
}
type LogUpdateRes struct {
	Id string `json:"id"`
}

type LogDeleteReq struct {
	g.Meta `path:"/log/delete" method:"post" tags:"LogService"`
	OrmIdReq
}
type LogDeleteRes struct {
	Data sql.Result
}

type LogCreateReq struct {
	g.Meta `path:"/log/create" method:"post" tags:"LogService"`
	LogAttrReq
}

type LogCreateRes struct {
	OrmIdReq
}

type LogAttrReq struct {
	UserId    int         `json:"userId"  dc:"用户id"  `    //
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`    // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间"`    // 修改时间
	DeletedAt *gtime.Time `json:"deletedAt" dc:"删除时间"`    // 删除时间
	IsDeleted int         `json:"isDeleted" dc:"数据的逻辑删除"` // 数据的逻辑删除
	Status    int         `json:"status"  dc:"状态"  `      // 状态
	Extra     string      `json:"extra"  dc:"额外信息"   `    //
	Ip        string      `json:"ip"     dc:"请求ip"   `    // 请求ip
	Path      string      `json:"path"    dc:"请求路径"  `    // 请求路径
	Method    string      `json:"method"   dc:"请求方法" `    // 请求方法
	Request   string      `json:"request"  dc:"请求参数" `    // 请求参数
	Response  string      `json:"response" dc:"响应内容" `    // 响应内容
	Agent     string      `json:"agent"    dc:"代理" `      // 代理
	Latency   int         `json:"latency"   dc:"延迟"`      // 延迟
	ErrMsg    string      `json:"errMsg"    dc:"错误信息"`    // 错误信息
	UserName  string      `json:"userName"  dc:"用户名"`     // 用户
}
