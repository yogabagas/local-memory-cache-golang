package service

import (
	"time"

	"github.com/koding/cache"
	kcache "github.com/koding/cache"
)

type CacheConn struct {
	cache *kcache.MemoryTTL
}

func NewConnKCache() *CacheConn {
	c := kcache.NewMemoryWithTTL(time.Duration(1 * time.Second))

	return &CacheConn{cache: c}
}

func (c *CacheConn) Set(key, value interface{}) error {
	err := c.cache.Set(key.(string), value)
	if err != nil {
		return err
	}
	return nil
}

func (c *CacheConn) Get(key interface{}) (interface{}, error) {
	resp, err := c.cache.Get(key.(string))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type CacheSharedConn struct {
	shared *kcache.ShardedTTL
}

func NewShardedCacheWithTTL() *CacheSharedConn {
	s := kcache.NewShardedCacheWithTTL(5*time.Millisecond, func() kcache.Cache {
		return cache.NewLRU(1000000)
	})
	return &CacheSharedConn{shared: s}
}

func (s *CacheSharedConn) Set(partition string, key, value interface{}) error {
	if err := s.shared.Set(partition, key.(string), value); err != nil {
		return err
	}
	return nil
}

func (s *CacheSharedConn) Get(partition string, key interface{}) (interface{}, error) {
	resp, err := s.shared.Get(partition, key.(string))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
