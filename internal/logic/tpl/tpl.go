package tpl

import (
	"admin/internal/service"
	"context"
	"fmt"
)

type sTpl struct {
	avatarUploadPath      string // 头像上传路径
	avatarUploadUrlPrefix string // 头像上传对应的URL前缀
}

func init() {
	instance := New()
	fmt.Println(instance)
	service.RegisterTpl(instance)
}

func New() *sTpl {
	return &sTpl{}
}

// Logout 注销
func (s *sTpl) Logout(ctx context.Context) error {
	//return service.Session().RemoveUser(ctx)
	return nil
}
