package model

// tpl 生成token
type TokenServiceGenTokenReq struct {
	// 用户id
	Uid string `json:"uid,omitempty"`
	// 过期时间戳，单位s
	Exp int64 `json:"exp,omitempty"`
}

type TokenServiceCheckTokenReq struct {
	Token string `json:"tokens"`
}
