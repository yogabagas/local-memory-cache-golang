package service

import (
	"time"

	tcache "github.com/ReneKroon/ttlcache"
)

type TCacheConn struct {
	cache *tcache.Cache
}

func NewClientConn() *TCacheConn {
	c := tcache.NewCache()
	c.SetCacheSizeLimit(3)
	c.SetTTL(15 * time.Second)
	return &TCacheConn{cache: c}
}

func (t *TCacheConn) Set(key, value interface{}) error {
	if err := t.cache.Set(key.(string), value); err != nil {
		return err
	}
	return nil
}

func (t *TCacheConn) Get(key interface{}) (interface{}, error) {
	resp, err := t.cache.Get(key.(string))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
