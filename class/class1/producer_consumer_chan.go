package main

import "fmt"
import "runtime"
import "sync/atomic"
import "errors"

const (
	NumItems     = 10
	NumConsumers = 4
)

type Buffer struct {
	fifo chan interface{}
}

func NewBuffer(n int) *Buffer {
	b := &Buffer{}
	b.fifo = make(chan interface{}, n)
	return b
}

func (b *Buffer) Free() {
	close(b.fifo)
}

func (b *Buffer) Add(item interface{}) {
	b.fifo <- item
}

func (b *Buffer) Get() (item interface{}, err error) {
	var ok bool
	item, ok = <-b.fifo
	if !ok {
		err = errors.New("closed")
	}
	return
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	buf := NewBuffer(NumItems / 4)

	// producer
	go func() {
		for i := 0; i < NumItems; i++ {
			buf.Add(i)
			fmt.Printf("Added %v\n", i)
		}
		buf.Free()
	}()

	// consumer
	var numRecv int32
	for cId := 0; cId < NumConsumers; cId++ {
		go func(id int) {
			for numRecv < NumItems {
				item, err := buf.Get()
				if err != nil {
					break
				}
				fmt.Printf("%v Got %v\n", id, item)
				atomic.AddInt32(&numRecv, 1)
			}
			fmt.Printf("%v done\n", id)
		}(cId)
	}

	var s string
	// wait for <enter> to exit
	fmt.Scanf("%s", &s)
}
