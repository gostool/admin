package v1

import "github.com/gogf/gf/v2/frame/g"

type ToolsCaptchaReq struct {
	g.Meta `path:"/captcha" method:"get"`
}
type ToolsCaptchaRes struct {
	Id   string `json:"captchaId" dc:"Random id"`
	Path string `json:"picPath" dc:"base64 img"`
}
