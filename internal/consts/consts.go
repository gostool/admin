package consts

import "errors"

const (
	DELETED      = 1
	CREATED      = 0
	ErrCodeOk    = 200
	ErrCodeBad   = 400
	ErrCodeToken = 403
	ErrCodeErr   = 500
)

const (
	LoggerDebug = "debug"
	ErrMsgErr   = "服务器出错"
)

var (
	ErrDel     = errors.New("删除失败")
	ErrUpdate  = errors.New("更新失败")
	ErrNotExit = errors.New("不存在")
	//CtxUserId common user info
	CtxUserId       = "ctx_user_id"
	CtxUserRoleId   = "ctx_user_role_id"
	CtxUserPassport = "ctx_user_passport"
	CtxUserNickName = "ctx_user_nickname"
	CtxUserRoleIds  = "ctx_user_roleIds"
	CtxUserData     = "ctx_user_data"
	CtxUserToken    = "ctx_user_token"
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
