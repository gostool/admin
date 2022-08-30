package model

import (
	"admin/internal/consts"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type MenuListInput struct {
	Page     int
	PageSize int
}

type MenuCreateInput struct {
	MenuAttr
}

type MenuDetailInput struct {
	Id int // role id
}

type MenuUpdateInput struct {
	Id int
	MenuAttr
}

type MenuDeleteInput struct {
	Id int
}

type MenuAttr struct {
	Title       string `json:"title"       ` // 名称
	Icon        string `json:"icon"        ` // 图标
	KeepAlive   int    `json:"keepAlive"   ` // 是否缓存
	DefaultMenu int    `json:"defaultMenu" ` // 是否基础路由
	Hidden      int    `json:"hidden"      ` //
	Sort        int    `json:"sort"        ` //
	ParentId    int    `json:"parentId"    ` //
	Name        string `json:"name"        ` //
	Path        string `json:"path"        ` //
	Component   string `json:"component"   ` //
}

func (r *MenuCreateInput) ToMap() (data g.Map) {
	data = g.Map{
		"title":        r.Title,
		"icon":         r.Icon,
		"keep_alive":   r.KeepAlive,
		"default_menu": r.DefaultMenu,
		"hidden":       r.Hidden,
		"sort":         r.Sort,
		"name":         r.Name,
		"path":         r.Path,
		"component":    r.Component,
		"is_deleted":   consts.CREATED,
	}
	if r.ParentId != -1 {
		data["parent_id"] = r.ParentId
	}
	return data
}

func (r *MenuUpdateInput) ToWhereMap() (data g.Map) {
	data = g.Map{
		"id":         r.Id,
		"is_deleted": consts.CREATED,
	}
	return data
}

func (r *MenuUpdateInput) ToMap() (data g.Map) {
	data = g.Map{}
	if r.ParentId != 0 {
		data["parent_id"] = r.ParentId
	}
	if r.ParentId == -1 {
		data["parent_id"] = gdb.Raw("NULL")
	}
	if r.Name != "" {
		data["name"] = r.Name
	}
	if r.Path != "" {
		data["path"] = r.Path
	}
	if r.Title != "" {
		data["title"] = r.Title
	}
	if r.Icon != "" {
		data["icon"] = r.Icon
	}
	data["keep_alive"] = r.KeepAlive
	data["default_menu"] = r.DefaultMenu
	data["hidden"] = r.Hidden
	data["sort"] = r.Sort
	return data
}
