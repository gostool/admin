package enforcer

import (
	"admin/internal/consts"
	"admin/internal/logic/adapter"
	"admin/internal/model"
	"admin/internal/model/serializer"
	"admin/internal/service"
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

type sEnforcer struct {
	enforcer *casbin.SyncedEnforcer
}

var logger *glog.Logger

func init() {
	logger = g.Log(consts.LoggerDebug)
	instance := New()
	service.RegisterEnforcer(instance)
}

func New() *sEnforcer {
	v := g.Cfg().MustGet(context.TODO(), "app")
	conf := v.MapStrVar()
	casbinConf := conf["casbin"].MapStrVar()
	model := casbinConf["model"].String()
	adapterInstance := adapter.New()
	enforcer, err := casbin.NewSyncedEnforcer(model, adapterInstance)
	enforcer.StartAutoLoadPolicy(5 * time.Second)
	if err != nil {
		panic(err)
	}
	return &sEnforcer{
		enforcer: enforcer,
	}
}

func (s *sEnforcer) Enforcer(ctx context.Context) (enforcer *casbin.SyncedEnforcer) {
	return s.enforcer
}

func (s *sEnforcer) List(ctx context.Context, in model.EnforcerListInput) (items []*serializer.CasbinApi, err error) {
	roleIdStr := gconv.String(in.RoleId)
	policyList := s.enforcer.GetFilteredPolicy(0, roleIdStr)
	items = make([]*serializer.CasbinApi, 0, len(policyList))
	for _, v := range policyList {
		items = append(items, &serializer.CasbinApi{
			Path:   v[1],
			Method: v[2],
		})
	}
	return items, nil
}

// Update 执行更新
func (s *sEnforcer) Update(ctx context.Context, in model.EnforcerUpdateInput) (err error) {
	roleIdStr := gconv.String(in.RoleId)
	s.Clear(0, roleIdStr)
	rules := make([][]string, 0, len(in.ApiInfoList))
	for _, v := range in.ApiInfoList {
		rules = append(rules, []string{roleIdStr, v.Path, v.Method})
	}
	g.Dump("rules:", rules)
	success, _ := s.enforcer.AddPolicies(rules)
	if !success {
		return consts.ErrUpdate
	}
	return nil
}

func (s *sEnforcer) Clear(v int, p ...string) bool {
	success, _ := s.enforcer.RemoveFilteredPolicy(v, p...)
	return success
}
