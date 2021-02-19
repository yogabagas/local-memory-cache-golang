package service

import (
	"time"

	cc "github.com/karlseguin/ccache"
)

type CcacheConn struct {
	cache *cc.Cache
}

func NewCCConn() *CcacheConn {
	c := cc.New(cc.Configure().MaxSize(3).GetsPerPromote(5))
	return &CcacheConn{cache: c}
}

func (c *CcacheConn) Get(key interface{}) interface{} {
	var resp interface{}

	itemFetch, err := c.cache.Fetch(key.(string), 5, func() (interface{}, error) {
		return resp, nil
	})
	if err == nil {
		return itemFetch.Value()
	}

	return nil
}

func (c *CcacheConn) Set(key, value interface{}) {
	c.cache.Set(key.(string), value, time.Duration(5*time.Second))
}
