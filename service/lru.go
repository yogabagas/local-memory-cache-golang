package service

import (
	lru "github.com/hashicorp/golang-lru"
)

type LRUConn struct {
	lru *lru.ARCCache
}

func NewLRUConnection() *LRUConn {
	arc, err := lru.NewARC(1000000)
	if err != nil {
		return nil
	}

	return &LRUConn{lru: arc}
}

func (l *LRUConn) Get(key interface{}) interface{} {
	if v, ok := l.lru.Get(key); ok {
		return v
	}
	return nil
}

func (l *LRUConn) Set(key interface{}, value interface{}) {
	l.lru.Add(key, value)
}

func (l *LRUConn) Remove(key string) {
	l.lru.Remove(key)
}

func (l *LRUConn) Length() int {
	return l.lru.Len()
}

func (l *LRUConn) Keys() []interface{} {
	return l.lru.Keys()
}

func (l *LRUConn) BulkSet(key interface{}, values []interface{}) {
	keyMap := make(map[interface{}]interface{})
	for _, value := range values {
		keyMap[key] = value
	}

	if v, ok := keyMap[key]; ok {
		l.lru.Add(key, v)
	}
}

// func (l *LRUConn) BulkGet(key interface{})  {
// 	if v, ok := l.lru.Get(key); ok {
// 		*val = append(*val, v)
// 	}
// }
