package v1

import (
	"admin/internal/model/serializer"
	"database/sql"
	"github.com/gogf/gf/v2/frame/g"
)

type MenuListReq struct {
	g.Meta `path:"/menu/list" method:"get" tags:"MenuService"`
	PageReq
}
type MenuListRes struct {
	Count int                `json:"count" dc:"记录总数"`
	Items []*serializer.Menu `json:"items" dc:"条目"`
}

type MenuDetailReq struct {
	g.Meta `path:"/menu/detail" method:"post" tags:"MenuService"`
	OrmIdReq
}
type MenuDetailRes struct {
	*serializer.Menu
}

type MenuUpdateReq struct {
	g.Meta `path:"/menu/update" method:"post" tags:"MenuService"`
	OrmIdReq
	MenuAttrReq
}
type MenuUpdateRes struct {
	Id string `json:"name"`
}

type MenuDeleteReq struct {
	g.Meta `path:"/menu/delete" method:"post" tags:"MenuService"`
	OrmIdReq
}
type MenuDeleteRes struct {
	Data sql.Result
}

type MenuCreateReq struct {
	g.Meta `path:"/menu/create" method:"post" tags:"MenuService"`
	MenuAttrReq
}

type MenuCreateRes struct {
	OrmIdReq
}

type MenuAttrReq struct {
	Title       string `json:"title"  dc:"显示名称"   `      // 名称
	Icon        string `json:"icon"   dc:"图标"     `      // 图标
	KeepAlive   int    `json:"keepAlive"  dc:"是否缓存" `    // 是否缓存
	DefaultMenu int    `json:"defaultMenu" dc:"是否基础路由" ` // 是否基础路由
	Hidden      int    `json:"hidden"     dc:"是否隐藏" `    //
	Sort        int    `json:"sort"       dc:"排序" `      //
	ParentId    int    `json:"parentId"   dc:"父ID" `     //
	Name        string `json:"name"       dc:"名字" `      //
	Path        string `json:"path"       dc:"路径" `      //
	Component   string `json:"component"  dc:"组件路径" `    //
}
