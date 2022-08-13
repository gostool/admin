package model

import "github.com/gogf/gf/v2/frame/g"

var UserFields = g.ArrayStr{
	"id",
	"name",
	"created_at",
}

var RoleFields = g.ArrayStr{
	"id",
	"name",
	"pid",
	"router",
	"created_at",
	"updated_at",
}

var MenuFields = g.ArrayStr{
	"id",
	"created_at",
	"updated_at",
	"title",
	"icon",
	"keep_alive",
	"default_menu",
	"hidden",
	"sort",
	"parent_id",
	"name",
	"path",
	"component",
}
