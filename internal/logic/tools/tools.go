package tools

import (
	"admin/internal/consts"
	"admin/internal/service"
	"admin/library/store"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/mojocn/base64Captcha"
)

var (
	captchaStore  *store.Cache
	captchaDriver = newDriver()
	logger        *glog.Logger
)

func init() {
	captchaStore = store.NewStoreCache(g.Redis(), 0)
	logger = g.Log(consts.LoggerDebug)
	instance := New()
	service.RegisterTools(instance)
}

type sTools struct{}

func New() *sTools {
	return &sTools{}
}

func newDriver() *base64Captcha.DriverString {
	driver := &base64Captcha.DriverString{
		Height:          44,
		Width:           126,
		NoiseCount:      5,
		ShowLineOptions: base64Captcha.OptionShowSineLine | base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowHollowLine,
		Length:          4,
		Source:          "1234567890",
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	return driver.ConvertFonts()
}

func (s *sTools) Captcha(ctx context.Context) (id, b64s string, err error) {
	var captcha = base64Captcha.NewCaptcha(captchaDriver, captchaStore)
	return captcha.Generate()
}

func (s *sTools) Verify(ctx context.Context, captchaId, captcha string, clear bool) (ok bool) {
	return captchaStore.Verify(captchaId, captcha, clear)
}
