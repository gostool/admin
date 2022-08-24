package model

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CasbinRuleListInput struct {
	Page     int
	PageSize int
}

type CasbinRuleCreateInput struct {
	CasbinRuleAttr
}

type CasbinRuleDetailInput struct {
	Id int // id
}

type CasbinRuleUpdateInput struct {
	Id int
	CasbinRuleAttr
}

type CasbinRuleDeleteInput struct {
	Id int
}

type CasbinRuleAttr struct {
	Id    uint64 `json:"id"    ` //
	Ptype string `json:"ptype" ` //
	V0    string `json:"v0"    ` //
	V1    string `json:"v1"    ` //
	V2    string `json:"v2"    ` //
	V3    string `json:"v3"    ` //
	V4    string `json:"v4"    ` //
	V5    string `json:"v5"    ` //
}

func (r *CasbinRuleCreateInput) ToMap() (data g.Map) {
	data = g.Map{
		"ptype": r.Ptype,
		"v0":    r.V0,
		"v1":    r.V1,
		"v2":    r.V2,
		"v3":    r.V3,
		"v4":    r.V4,
		"v5":    r.V5,
	}
	return data
}

func (r *CasbinRuleUpdateInput) ToWhereMap() (data g.Map) {
	data = g.Map{
		"id": r.Id,
	}
	return data
}

func (r *CasbinRuleUpdateInput) ToMap() (data g.Map) {
	data = g.Map{
		"id":    r.Id,
		"ptype": r.Ptype,
		"v0":    r.V0,
		"v1":    r.V1,
		"v2":    r.V2,
		"v3":    r.V3,
		"v4":    r.V4,
		"v5":    r.V5,
	}
	return data
}
