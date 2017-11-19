package avsgo

type Goroutine struct {
	value int32

	readCh  chan chan int32
	writeCh chan int32
}

func NewGoroutine() *Goroutine {
	return &Goroutine{
		readCh:  make(chan chan int32),
		writeCh: make(chan int32),
	}
}

func (a *Goroutine) Start() {
	for {
		select {
		case c := <-a.readCh:
			c <- a.value
		case v := <-a.writeCh:
			a.value = v
		}
	}
}

func (a *Goroutine) Read() int32 {
	c := make(chan int32)
	a.readCh <- c
	return <-c
}

func (a *Goroutine) Write(v int32) {
	a.writeCh <- v
}
