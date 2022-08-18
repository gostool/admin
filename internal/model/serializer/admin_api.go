package serializer

import "github.com/gogf/gf/v2/os/gtime"

type Api struct {
	Id        int         `json:"id"        ` //
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"-" `         // 修改时间
	DeletedAt *gtime.Time `json:"-" `         // 删除时间
	IsDeleted int         `json:"-" `         // 数据的逻辑删除
	Group     string      `json:"group"     ` // 分组
	Path      string      `json:"path"      ` // 请求路径
	Detail    string      `json:"detail"    ` // 简介
	Method    string      `json:"method"    ` // 请求方式
}
