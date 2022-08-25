package adapter

import (
	"admin/internal/consts"
	"admin/internal/dao"
	dbModel "admin/internal/model"
	"admin/internal/model/entity"
	"admin/internal/service"
	"context"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

type sAdapter struct {
}

var logger *glog.Logger

func init() {
	logger = g.Log(consts.LoggerDebug)
	instance := New()
	service.RegisterAdapter(instance)
}

func New() *sAdapter {
	return &sAdapter{}
}

func loadPolicyLine(line *entity.CasbinRule, model model.Model) {
	lineText := line.Ptype
	if line.V0 != "" {
		lineText += ", " + line.V0
	}
	if line.V1 != "" {
		lineText += ", " + line.V1
	}
	if line.V2 != "" {
		lineText += ", " + line.V2
	}
	if line.V3 != "" {
		lineText += ", " + line.V3
	}
	if line.V4 != "" {
		lineText += ", " + line.V4
	}
	if line.V5 != "" {
		lineText += ", " + line.V5
	}
	persist.LoadPolicyLine(lineText, model)
}

// LoadPolicy loads policy from database.
func (a *sAdapter) LoadPolicy(model model.Model) error {
	var lines []*entity.CasbinRule
	ctx := context.TODO()
	if err := dao.CasbinRule.Ctx(ctx).Scan(&lines); err != nil {
		return err
	}
	for _, line := range lines {
		loadPolicyLine(line, model)
	}
	return nil
}

func (a *sAdapter) dropTable() (err error) {
	return
}

func (a *sAdapter) createTable() (err error) {
	return
}

func (a *sAdapter) savePolicyLine(ptype string, line []string) (err error) {
	rule := convLineToRule(ptype, line)
	var casbinRuleCreateInput *dbModel.CasbinRuleCreateInput
	casbinRuleCreateInput.Ptype = rule.Ptype
	casbinRuleCreateInput.V0 = rule.V0
	casbinRuleCreateInput.V1 = rule.V1
	casbinRuleCreateInput.V2 = rule.V2
	casbinRuleCreateInput.V3 = rule.V3
	casbinRuleCreateInput.V4 = rule.V4
	casbinRuleCreateInput.V5 = rule.V5
	return a.savePolicyRule(casbinRuleCreateInput)
}

func (a *sAdapter) savePolicyRule(rule *dbModel.CasbinRuleCreateInput) (err error) {
	ctx := context.TODO()
	_, err = service.CasbinRule().Create(ctx, rule)
	if err != nil {
		return err
	}
	return nil
}

func convLineToRule(ptype string, rule []string) (line *entity.CasbinRule) {
	line = &entity.CasbinRule{}
	line.Ptype = ptype
	if len(rule) > 0 {
		line.V0 = rule[0]
	}
	if len(rule) > 1 {
		line.V1 = rule[1]
	}
	if len(rule) > 2 {
		line.V2 = rule[2]
	}
	if len(rule) > 3 {
		line.V3 = rule[3]
	}
	if len(rule) > 4 {
		line.V4 = rule[4]
	}
	if len(rule) > 5 {
		line.V5 = rule[5]
	}
	return line
}

func (a *sAdapter) delPolicyRule(rule *entity.CasbinRule) (err error) {
	ctx := context.TODO()
	_, err = service.CasbinRule().DeleteByModel(ctx, rule)
	return err
}

// SavePolicy saves policy to database.
func (a *sAdapter) SavePolicy(model model.Model) (err error) {
	err = a.dropTable()
	if err != nil {
		return
	}
	err = a.createTable()
	if err != nil {
		return
	}
	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			err = a.savePolicyLine(ptype, rule)
			if err != nil {
				return err
			}
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			err = a.savePolicyLine(ptype, rule)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// AddPolicy adds a policy rule to the storage.
func (a *sAdapter) AddPolicy(sec string, ptype string, rule []string) (err error) {
	err = a.savePolicyLine(ptype, rule)
	return err
}

func (a *sAdapter) delPolicyLine(ptype string, line []string) (err error) {
	rule := convLineToRule(ptype, line)
	return a.delPolicyRule(rule)
}

// RemovePolicy removes a policy rule from the storage.
func (a *sAdapter) RemovePolicy(sec string, ptype string, line []string) error {
	return a.delPolicyLine(ptype, line)
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *sAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	rule := &entity.CasbinRule{}
	rule.Ptype = ptype
	if fieldIndex <= 0 && 0 < fieldIndex+len(fieldValues) {
		rule.V0 = fieldValues[0-fieldIndex]
	}
	if fieldIndex <= 1 && 1 < fieldIndex+len(fieldValues) {
		rule.V1 = fieldValues[1-fieldIndex]
	}
	if fieldIndex <= 2 && 2 < fieldIndex+len(fieldValues) {
		rule.V2 = fieldValues[2-fieldIndex]
	}
	if fieldIndex <= 3 && 3 < fieldIndex+len(fieldValues) {
		rule.V3 = fieldValues[3-fieldIndex]
	}
	if fieldIndex <= 4 && 4 < fieldIndex+len(fieldValues) {
		rule.V4 = fieldValues[4-fieldIndex]
	}
	if fieldIndex <= 5 && 5 < fieldIndex+len(fieldValues) {
		rule.V5 = fieldValues[5-fieldIndex]
	}
	err := a.delPolicyRule(rule)
	return err
}
