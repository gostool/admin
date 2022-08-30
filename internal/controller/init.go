package controller

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
)

var logger *glog.Logger

func init() {
	logger = g.Log()
}

func ToRoleIdsNums(roleIds string) []int {
	return gconv.Ints(roleIds)
}
