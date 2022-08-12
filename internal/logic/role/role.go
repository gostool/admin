package role

import (
	"admin/internal/service"
	"context"
	"fmt"
)

type sRole struct {
}

func init() {
	instance := New()
	fmt.Println(instance)
	service.RegisterRole(instance)
}

func New() *sRole {
	return &sRole{}
}

func (s *sRole) List(ctx context.Context) error {
	//return service.Session().RemoveUser(ctx)
	return nil
}
