package avsgo

import (
	"testing"
	"time"
)

type Writer interface {
	Write(int32)
}

func BenchmarkGoroutine(b *testing.B) {
	b.StopTimer()
	a := NewGoroutine()
	go a.Start()
	setupTickWriter(a)
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		a.Read()
	}
}

func BenchmarkAtomic(b *testing.B) {
	b.StopTimer()
	a := NewAtomic()
	setupTickWriter(a)
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		a.Read()
	}
}

func BenchmarkMutex(b *testing.B) {
	b.StopTimer()
	a := NewMutex()
	setupTickWriter(a)
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		a.Read()
	}
}

func setupTickWriter(a Writer) {
	ticker := time.NewTicker(10 * time.Millisecond)
	go func() {
		for range ticker.C {
			a.Write(123)
		}
	}()
}
