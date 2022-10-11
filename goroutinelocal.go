package goroutinelocal

import (
	"github.com/sonnt85/gosystem"
	"sync"
)

type goroutineLocal[T any] struct {
	initfun func() T
	m       *sync.Map
}

func NewGoroutineLocal[T any](initfun func() T) *goroutineLocal[T] {
	return &goroutineLocal[T]{initfun: initfun, m: &sync.Map{}}
}

func (gl *goroutineLocal[T]) Get() T {
	value, ok := gl.m.Load(gosystem.GetGoroutineId())
	if !ok && gl.initfun != nil {
		value = gl.initfun()
	}
	ret, _ := value.(T)
	return ret
}

func (gl *goroutineLocal[T]) GetMap() (m map[int64]T) {
	m = make(map[int64]T, 0)
	gl.m.Range(func(key, value any) bool {
		k, _ := key.(int64)
		v, _ := value.(T)
		m[k] = v
		return true
	})
	return
}

func (gl *goroutineLocal[T]) Set(v T) {
	gl.m.Store(gosystem.GetGoroutineId(), v)
}

func (gl *goroutineLocal[T]) Remove() {
	gl.m.Delete(gosystem.GetGoroutineId())
}
