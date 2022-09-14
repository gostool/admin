package tests

import (
	"admin/library/store"
	"fmt"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"testing"
	"time"
)

func init() {
	ctx = gctx.New()
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("manifest/config/config.yaml")
}

func TestGfRedis(t *testing.T) {
	v, _ := g.Redis().Do(ctx, "SET", "k", "v")
	fmt.Println(v.String())
}

func TestExampleCache_SetAdapter(t *testing.T) {
	v := g.Cfg().MustGet(ctx, "redis")
	conf := v.MapStrVar()
	redisConf := conf["default"].Map()
	g.Dump(redisConf)
	var (
		err         error
		ctx         = gctx.New()
		cache       = gcache.New()
		redisConfig = &gredis.Config{
			Address: gconv.String(redisConf["address"]),
			Db:      gconv.Int(redisConf["db"]),
		}
		cacheKey   = `key`
		cacheValue = `value`
	)
	// Create redis client object.
	redis, err := gredis.New(redisConfig)
	if err != nil {
		panic(err)
	}
	// Create redis cache adapter and set it to cache object.
	cache.SetAdapter(gcache.NewAdapterRedis(redis))

	// Set and Get using cache object.
	err = cache.Set(ctx, cacheKey, cacheValue, time.Second*60)
	if err != nil {
		panic(err)
	}
	fmt.Println(cache.MustGet(ctx, cacheKey).String())

	// Get using redis client.
	fmt.Println(redis.MustDo(ctx, "GET", cacheKey).String())

	// Output:
	// value
	// value
}

func TestCacheRedis(t *testing.T) {
	redisConfig := &gredis.Config{
		Address: "127.0.0.1:6379",
		Db:      9,
	}
	redis, err := gredis.New(redisConfig)
	if err != nil {
		t.Fatal(err)
	}
	ttl := time.Second * 60 * 24
	s := store.NewStoreCache(redis, ttl)
	s.Set("1", "tp")
	s.Verify("1", "tp", false)
}
