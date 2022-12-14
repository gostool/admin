package model

import (
	"github.com/gogf/gf/v2/frame/g"
)

var UserFields = g.ArrayStr{
	"id",
	"passport",
	"created_at",
	"updated_at",
	"password",
	"nickname",
	"roleId",
	"roleIds",
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

var RoleMenuFields = g.ArrayStr{
	"id",
	"created_at",
	"updated_at",
	"role_id",
	"menu_id",
	"is_deleted",
}

var LogFields = g.ArrayStr{
	"id",
	"userId",
	"createdAt",
	"updatedAt",
	"deletedAt",
	"isDeleted",
	"status",
	"extra",
	"ip",
	"path",
	"method",
	"request",
	"response",
	"agent",
	"latency",
	"errMsg",
	"userName",
}

var ApiFields = g.ArrayStr{
	"id",
	"created_at",
	"updated_at",
	"deleted_at",
	"is_deleted",
	"group",
	"path",
	"detail",
	"method",
}

var CasbinRuleFields = g.ArrayStr{
	"id",
	"ptype",
	"v0",
	"v1",
	"v2",
	"v3",
	"v4",
	"v5",
}
