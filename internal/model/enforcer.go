package model

type EnforcerListInput struct {
	RoleId int
}

type EnforcerUpdateInput struct {
	RoleId      int
	ApiInfoList []AdminApiAttr
}
