package tests

import (
	"testing"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/test/gtest"
)

// 测试针对: config_tpl.toml 进行测试.
// 项目中真实使用.config.toml.
// config_tpl.toml 必须实时同步.
// 更新config_tpl.toml 需要更新本测试文件. git action会进行测试
// 本地测试: go test -v ./config

func TestCfgSetPath(t *testing.T) {
	g.Cfg().SetFileName("config_tpl.toml")
	v := g.Config().GetVar("server.Address")
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(v.String(), ":8299")
	})
}

func TestLogger(t *testing.T) {
	g.Cfg().SetFileName("config_tpl.toml")
	// 多个日志实例
	// 1.默认日志
	g.Log().Info("i am in logger default")
	g.Log("debug").Info("i am in logger debug")
	g.Log("test").Info("i am in logger test")
}

// add redis conf test
func TestCfgRedis(t *testing.T) {
	g.Cfg().SetFileName("config_tpl.toml")
	v := g.Config().GetVar("redis")
	redis := v.MapStrStr()
	gtest.C(t, func(t *gtest.T) {
		t.AssertEQ(redis["default"], "127.0.0.1:6379,0")
		t.AssertEQ(redis["cache"], "127.0.0.1:6379,1,123456?idleTimeout=600")
	})
}

func TestSnowAppConf(t *testing.T) {
	g.Cfg().SetFileName("manifest/config/config.yaml")
	v := g.Config().GetVar("app")
	conf := v.MapStrVar()
	t.Log(conf)
	gtest.C(t, func(t *gtest.T) {
		//t.AssertEQ(conf["shareLink"].String(), "http://192.168.0.100:9000/preview/#/index")
		//t.AssertEQ(conf["machineId"].Uint16(), uint16(1))
		//// jwt
		//t.AssertEQ(conf["jwtSalt"].String(), "testsalt")
		//t.AssertEQ(conf["jwtExp"].Int64(), int64(604800))
		//// shp
		//t.AssertEQ(conf["shpOutMinLevel"].Int(), -3)
		//t.AssertEQ(conf["shpOutMaxLevel"].Int(), 0)
		//t.AssertEQ(conf["shpInDoorMinLevel"].Int(), 1)
		//t.AssertEQ(conf["shpInDoorMaxLevel"].Int(), 10)
		//// theme
		//t.AssertEQ(conf["themeIconSysDir"].String(), "/static/mapimgs")
		//// op log size
		//t.AssertEQ(conf["opLogPoolSize"].Int(), 5)
		//// tmp dir delete
		//t.AssertEQ(conf["tmpDirDelete"].Bool(), true)
	})

}
