type Tpl struct {
	OrmCommon
	Name     string `json:"name"      ` //
	Password string `json:"password"  ` //
	Nickname string `json:"nickname"  ` //
	RoleId   int    `json:"roleId"    ` //
}

func (u *Tpl) ToData() (data *g.Map) {
	data = &g.Map{
		"id":       u.Id,
		"name":     u.Name,
		"nickname": u.Name,
	}
	return data
}