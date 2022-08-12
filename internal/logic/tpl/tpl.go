package tpl

import (
	"context"
	"fmt"
)

type sTpl struct {
}

func init() {
	instance := New()
	fmt.Println(instance)
	//service.RegisterTpl(instance)
}

func New() *sTpl {
	return &sTpl{}
}

// Logout 注销
func (s *sTpl) Logout(ctx context.Context) error {
	//return service.Session().RemoveUser(ctx)
	return nil
}
