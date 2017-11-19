package avsgo

type Goroutine struct {
	value int

	readCh  chan chan int
	writeCh chan int
}

func NewGoroutine() *Goroutine {
	return &Goroutine{
		readCh:  make(chan chan int),
		writeCh: make(chan int),
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

func (a *Goroutine) Read() int {
	c := make(chan int)
	a.readCh <- c
	return <-c
}

func (a *Goroutine) Write(v int) {
	a.writeCh <- v
}
