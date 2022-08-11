package v1

import "github.com/gogf/gf/v2/frame/g"

type ToolCaptchaReq struct {
	g.Meta `path:"/captcha" method:"get" tags:"ToolService" summary:"generate captcha"`
}

type ToolCaptchaRes struct {
	Id   string `json:"captchaId"`
	Path string `json:"picPath"`
}
