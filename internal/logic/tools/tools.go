package tools

import (
	"admin/internal/consts"
	"admin/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/mojocn/base64Captcha"
)

const (
	keyLong   = 6
	imgWidth  = 240
	imgHeight = 80
)

var Store = base64Captcha.DefaultMemStore

type sTools struct {
}

var logger *glog.Logger

func init() {
	logger = g.Log(consts.LoggerDebug)
	instance := New()
	service.RegisterTools(instance)
}

func New() *sTools {
	return &sTools{}
}

func (s *sTools) Captcha(ctx context.Context) (id, b64s string, err error) {
	var driver = base64Captcha.NewDriverDigit(imgHeight, imgWidth, keyLong, 0.7, 80) // 字符,公式,验证码配置, 生成默认数字的driver
	var captcha = base64Captcha.NewCaptcha(driver, Store)
	return captcha.Generate()
}

func (s *sTools) Verify(ctx context.Context, captchaId, captcha string, clear bool) (ok bool) {
	return Store.Verify(captchaId, captcha, clear)
}
