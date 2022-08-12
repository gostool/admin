package tpl

import (
	"admin/internal/service"
	"context"
)

type sTpl struct {
}

func init() {
	instance := New()
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
