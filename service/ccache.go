package service

import (
	"time"

	cc "github.com/karlseguin/ccache"
)

type CcacheConn struct {
	cache *cc.Cache
}

func NewCCConn() *CcacheConn {
	c := cc.New(cc.Configure().MaxSize(1000000).GetsPerPromote(5))
	return &CcacheConn{cache: c}
}

func (c *CcacheConn) Get(key interface{}) interface{} {
	item := c.cache.Get(key.(string))
	return item.Value()
}

func (c *CcacheConn) Set(key, value interface{}) {
	c.cache.Set(key.(string), value, time.Duration(2*time.Minute))
}
