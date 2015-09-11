package main

import "fmt"
import "container/list"
import "sync"
import "runtime"

const (
	NumItems = 20
)

type Buffer struct {
	capacity  int
	data      list.List
	mu        sync.Mutex
	emptyCond *sync.Cond
	fullCond  *sync.Cond
}

func NewBuffer(n int) *Buffer {
	b := &Buffer{}
	b.capacity = n
	b.mu = sync.Mutex{}
	b.emptyCond = sync.NewCond(&b.mu)
	b.fullCond = sync.NewCond(&b.mu)
	return b
}

func (b *Buffer) Add(item interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for b.data.Len() >= b.capacity {
		b.fullCond.Wait()
	}
	b.data.PushBack(item)
	b.emptyCond.Signal()
	return
}

func (b *Buffer) Get() (item interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for b.data.Len() == 0 {
		b.emptyCond.Wait()
	}
	item = b.data.Remove(b.data.Front())
	b.fullCond.Signal()
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
	}()

	// consumer
	go func() {
		count := 0
		for count < NumItems {
			item := buf.Get()
			fmt.Printf("Got %v\n", item)
			count++
		}
	}()

	var s string
	// wait for <enter> to exit
	fmt.Scanf("%s", &s)
}
