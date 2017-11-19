package avsgo

import (
	"sync"
)

type Mutex struct {
	l     sync.RWMutex
	value int32
}

func NewMutex() *Mutex {
	return &Mutex{}
}

func (a *Mutex) Read() int32 {
	a.l.RLock()
	defer a.l.RUnlock()
	return a.value
}

func (a *Mutex) Write(v int32) {
	a.l.Lock()
	defer a.l.Unlock()
	a.value = v
}
