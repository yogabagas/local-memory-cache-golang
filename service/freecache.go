package service

import (
	"bytes"
	gob "encoding/gob"

	f "github.com/coocood/freecache"
)

type FCacheConn struct {
	fcache *f.Cache
}

func NewFreeCacheConn() *FCacheConn {
	// size := 1000 * 1024 * 1024
	cache := f.NewCache(3)
	return &FCacheConn{fcache: cache}
}

func (fc *FCacheConn) Set(key, value interface{}) error {
	var buf bytes.Buffer
	gob.NewEncoder(&buf).Encode(value)
	if err := fc.fcache.Set([]byte(key.(string)), buf.Bytes(), 5); err != nil {
		return err
	}
	return nil
}

func (fc *FCacheConn) Get(key interface{}) (interface{}, error) {
	var res interface{}

	resp, err := fc.fcache.Get([]byte(key.(string)))
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(resp)
	gob.NewDecoder(buf).Decode(res)

	return res, nil
}
