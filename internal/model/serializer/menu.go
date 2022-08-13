package serializer

import "github.com/gogf/gf/v2/os/gtime"

type Menu struct {
	Id          int         `json:"id"          ` //
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 修改时间
	Title       string      `json:"title"       ` // 名称
	Icon        string      `json:"icon"        ` // 图标
	KeepAlive   int         `json:"keepAlive"   ` // 是否缓存
	DefaultMenu int         `json:"defaultMenu" ` // 是否基础路由
	Hidden      int         `json:"hidden"      ` //
	Sort        int         `json:"sort"        ` //
	ParentId    int         `json:"parentId"    ` //
	Name        string      `json:"name"        ` //
	Path        string      `json:"path"        ` //
	Component   string      `json:"component"   ` //
}
