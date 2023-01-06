package utils

import (
	"github.com/patrickmn/go-cache"
	"time"
)
type GoCache struct {
     cache  *cache.Cache
}

func NewCache() *GoCache {
	// 创建一个默认过期时间为5分钟的缓存适配器
	// 每60清除一次过期的项目
	go_cache := GoCache{}
	go_cache.cache = cache.New(5*time.Minute, 60*time.Second)
	return &go_cache
}

func (c *GoCache) SetCahce(k string, x interface{}, d time.Duration) {
	c.cache.Set(k, x, d)
}

func (c *GoCache) GetCache(k string) (interface{}, bool) {
	return c.cache.Get(k)
}


func (c *GoCache) SetDefaultCahce (k string, x interface{}) {
	c.cache.SetDefault(k, x)
}

func (c *GoCache) DeleteCache(k string) {
	c.cache.Delete(k)
}
