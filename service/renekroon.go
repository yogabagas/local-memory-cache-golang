package service

import (
	"time"

	tcache "github.com/ReneKroon/ttlcache/v2"
)

type TCacheConn struct {
	cache *tcache.Cache
}

func NewClientConn() *TCacheConn {
	c := tcache.NewCache()
	c.SetCacheSizeLimit(3)
	c.SetTTL(5 * time.Second)
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

func (t *TCacheConn) SetWithTTL(key, value interface{}) error {
	if err := t.cache.SetWithTTL(key.(string), value, 5); err != nil {
		return err
	}
	return nil
}

func (t *TCacheConn) SetWithGoRoutine(key, value interface{}) error {
	ch := make(chan error)
	go func() {
		err := t.cache.Set(key.(string), value)
		ch <- err
		defer func() {
			if r := recover(); r != nil {
				return
			}
		}()
	}()

	for {
		select {
		case err, ok := <-ch:
			if ok {
				return err
			}
			close(ch)
			return nil
		default:
			close(ch)
			return nil
		}
	}
}

func (t *TCacheConn) GetWithGoRoutine(key interface{}) (interface{}, error) {
	respCh := make(chan interface{})
	errCh := make(chan error)

	go func() {
		resp, err := t.cache.Get(key.(string))
		if err != nil {
			errCh <- err
		}
		respCh <- resp

		defer func() {
			if r := recover(); r != nil {
				return
			}
		}()
	}()

	for {
		select {
		case resp, ok := <-respCh:
			if ok {
				return resp, nil
			}
		case err, ok := <-errCh:
			if ok {
				return nil, err
			}
		case <-time.After(3 * time.Second):
			close(respCh)
			close(errCh)
		}
	}
	// return nil, nil
}
