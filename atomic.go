package avsgo

import "sync/atomic"

type Atomic struct {
	value int32
}

func NewAtomic() *Atomic {
	return &Atomic{}
}

func (a *Atomic) Read() int32 {
	return atomic.LoadInt32(&a.value)
}

func (a *Atomic) Write(v int32) {
	atomic.StoreInt32(&a.value, v)
}
