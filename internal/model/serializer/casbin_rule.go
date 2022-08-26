package serializer

type CasbinRule struct {
	Id    uint64 `json:"id"    ` //
	Ptype string `json:"ptype" ` //
	V0    string `json:"v0"    ` //
	V1    string `json:"v1"    ` //
	V2    string `json:"v2"    ` //
	V3    string `json:"v3"    ` //
	V4    string `json:"v4"    ` //
	V5    string `json:"v5"    ` ///
}

type CasbinApi struct {
	Path   string `json:"path"      ` // 请求路径
	Method string `json:"method"    ` // 请求方式
}
