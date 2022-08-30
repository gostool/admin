package v1

import (
	"admin/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type StatusReq struct {
	g.Meta `path:"/status/list" method:"get" tags:"ServiceStatus"`
}

type StatusRes struct {
	*model.Server
}
