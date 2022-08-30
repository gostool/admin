package tests

import (
	v1 "admin/api/v1"
	"admin/internal/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"testing"
)

func TestStruct(t *testing.T) {
	req := &v1.AdminCasbinUpdateReq{
		RoleId: 1,
		ApiInfoList: []v1.AdminCasbinAttr{
			{
				Path:   "data",
				Method: "read",
			},
		},
	}
	in := model.EnforcerUpdateInput{
		RoleId: req.RoleId,
	}
	infoList := make([]model.AdminCasbinAttr, len(req.ApiInfoList))
	err := gconv.Struct(req.ApiInfoList, &infoList)
	if err != nil {
		t.Fatal(err)
	}
	in.ApiInfoList = infoList
	g.Dump(in)
}
