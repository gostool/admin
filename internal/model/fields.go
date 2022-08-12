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
