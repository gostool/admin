package store

import (
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"time"
)

const ExpireTime = time.Second * 60 * 5

type Cache struct {
	cache *gcache.Cache
	ttl   time.Duration
}

func NewStoreCache(redis *gredis.Redis, ttl time.Duration) *Cache {
	cache := gcache.New()
	cache.SetAdapter(gcache.NewAdapterRedis(redis))
	if ttl <= 0 {
		ttl = ExpireTime
	}
	return &Cache{
		cache: cache,
		ttl:   ttl,
	}
}
func (s Cache) Set(id string, value string) {
	ctx := gctx.New()
	s.cache.Set(ctx, id, value, s.ttl)
}

func (s Cache) Get(id string, clear bool) string {
	ctx := gctx.New()
	v := s.cache.MustGet(ctx, id).String()
	if clear {
		s.cache.Remove(ctx, id)
	}
	return v
}

func (s Cache) Verify(id, answer string, clear bool) bool {
	return s.Get(id, clear) == answer
}
