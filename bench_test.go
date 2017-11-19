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

func BenchmarkGoroutineParallel(b *testing.B) {
	b.StopTimer()
	a := NewGoroutine()
	go a.Start()
	setupTickWriter(a)
	b.StartTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			a.Read()
		}
	})
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

func BenchmarkAtomicParallel(b *testing.B) {
	b.StopTimer()
	a := NewAtomic()
	setupTickWriter(a)
	b.StartTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			a.Read()
		}
	})
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

func BenchmarkMutexParallel(b *testing.B) {
	b.StopTimer()
	a := NewMutex()
	setupTickWriter(a)
	b.StartTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			a.Read()
		}
	})
}

func setupTickWriter(a Writer) {
	ticker := time.NewTicker(10 * time.Millisecond)
	go func() {
		for range ticker.C {
			a.Write(123)
		}
	}()
}
