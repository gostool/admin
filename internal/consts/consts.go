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

func OrderFiledByType(c int) string {
	switch c {
	case 0:
		return "created_at desc"
	case 1:
		return "created_at asc"
	case 2:
		return "updated_at desc"
	case 3:
		return "updated_at asc"
	case -1:
		return "id desc"
	default:
		return "created_at desc"
	}
}
