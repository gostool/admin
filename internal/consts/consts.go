package consts

import "errors"

const (
	DELETED     = 1
	CREATED     = 0
	LoggerDebug = "debug"
)

var (
	ErrDel     = errors.New("删除失败")
	ErrUpdate  = errors.New("更新失败")
	ErrNotExit = errors.New("不存在")
)
