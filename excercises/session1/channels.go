package main

import (
	"fmt"
	"time"
)

const (
	Num = 5
)

func channelProcessor(name string, c chan int) {
	for {
		var i int

		// ok is true if the channel is still open
		i, ok := <-c
		if !ok {
			fmt.Printf("%v closed\n", name)
			break
		} else {
			fmt.Printf("%v: %v\n", name, i)
		}
	}
}

func main() {
	// create the channels
	var unbufferedChan chan int = make(chan int)
	var bufferedChan chan int = make(chan int, Num)

	// launch the receive goroutines in the background
	go channelProcessor("unbuffered", unbufferedChan)
	go channelProcessor("buffered", bufferedChan)

	// send values to the channels
	go func() {
		for i := 0; i < Num; i++ {
			fmt.Printf("send %v\n", i)
			unbufferedChan <- i
			bufferedChan <- i
		}
	}()

	// give the goroutines a chance to run
	time.Sleep(5 * time.Second)
	close(bufferedChan)
	close(unbufferedChan)
	time.Sleep(1 * time.Second)
}
