package avsgo

import (
	"testing"
	"time"
)

func BenchmarkGoroutine(b *testing.B) {
	b.StopTimer()
	a := NewGoroutine()
	go a.Start()

	ticker := time.NewTicker(10 * time.Millisecond)
	go func() {
		for range ticker.C {
			a.Write(123)
		}
	}()
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		a.Read()
	}
}

func BenchmarkAtomic(b *testing.B) {
	b.StopTimer()
	a := NewAtomic()
	ticker := time.NewTicker(10 * time.Millisecond)
	go func() {
		for range ticker.C {
			a.Write(123)
		}
	}()
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		a.Read()
	}
}

func BenchmarkMutex(b *testing.B) {
	b.StopTimer()
	a := NewMutex()
	ticker := time.NewTicker(10 * time.Millisecond)
	go func() {
		for range ticker.C {
			a.Write(123)
		}
	}()
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		a.Read()
	}
}
