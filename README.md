# CSCI-GA.3033.2015
Course Website: http://news.cs.nyu.edu/~jinyang/fa15-ds/

## Part 1 - Golang Basics

This document assumes you have already reviewed [A Tour of Go] [1]. After you have completed the tour, bookmark [Effective Go] [2] as it will be extremely useful during your project development. The tutorial [How to Write Go Code] [5] contains information about how the primary tools in the Go ecosystem are configured.

Go is a staticly typed, garbage collected language that is heavily influenced by C 
-- and to a lesser extent C++ and Java. It's type system contains similar basic types 
as the aforementioned languages such as int, float, double, string, structs, 
functions, arrays, and pointers. Here is a program demonstrating how to declare these basic types. 

	package main
	
	import "fmt"
	
	type Record struct {
		id   int
		name string
	}
	
	func main() {
		var s string = "foo"
		var f float64 = 2.222
	
		var i int = 123
		// take address of i
		var ptrToInt *int = &i
		// dereference pointer
		var i2 = *ptrToInt

		var intAry = [3]int {1, 2, 3}
	
		var myStruct Record
		myStruct.id = 999
		myStruct.name = "my record"
	
		// alternate initialization based on field definition positon
		var myStruct2 Record = Record{1000, "my other record"}
	
		fmt.Printf("%v %v %v %v %v %v %v %v\n", s, f, i, ptrToInt, intAry, i2, myStruct, myStruct2)
	}

We can omit the type name in variable declarations using this alternate syntax:

	x := 123
	s := "golang string"

This document will use the more verbose syntax for clarity.


Golang also offers other primitive types that C++ and Java do not offer such as channels, slices, and maps. We cover channels and slices in this document.

### Slices
[Slices] [3] are a type that can be considered a lightweight array. More accurately, it is an abstraction of a view into an array that contains a pointer to the array, a length, and a capacity. Since functions in Go have call-by-value semantics, for larger data sets we can pass a slice more efficiently than an actual array. This also allows one to acheive the same effect as call-by-reference, since modifying an element of a slice will allow changes to the backing array to be visible to the caller.

We can create a slice by taking a slice of an existing array or creating a new slice and associated backing array) in one step. The following code shows how to create slices and assign elements:

	package main
	
	import "fmt"
	
	func main() {
		var intAry = [4]int{1, 2, 3, 4}
	
		// creates a new slice; does not allocate a new backing array
		var intSlice = intAry[1:3] // contians elements [2,3]
	
		// changes to the slice are reflected in the array and vice-versa
		intSlice[1] = -1
		intAry[1] = -2
	
		// creates a new slice while also allocating storage for a backing array
		var intSlice2 = make([]int, 4, 4) // contains [0 0 0 0]
		intSlice2[0] = -3
		intSlice2[1] = -2
		intSlice2[2] = -1

		// output: [1 -2 -1 4] [-2 -1] [-3 -2 -1 0]
		fmt.Printf("elements: %v %v %v\n", intAry, intSlice, intSlice2)
	
		// get the length of a slice (or array) with the 'len' builtin
		fmt.Printf("lengths: %v %v %v\n", len(intAry), len(intSlice), len(intSlice2))
	}

In general, using slices instead of array types is considered [effective go code] [2].

### Channels

[Channels] [4] are an FIFO queue abstraction that allows multiple goroutines to communicate. There are two types: buffered and unbuffered. An unbuffered channel will block the sender until the value is received; while the buffered channel will only block the sender if the buffer is full. Conversely, the receiver will be blocked if there is no data in the channel. This blocking behavior can be exploited to provide similar concurrency control mechanisms as a traditional Mutex. A common pattern is to launch multiple goroutines for a long running task, and have each goroutine report it's results via a shared channel. We show how to create, send, and receive data on a channel below:


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

### Part 2 - Excercises
We will look at your solutions the excercises below during the class. 

- Complete the excercises in sum.go and duplicate.go in this repository. For the sum.go problem, try to implement using channels and without. Which one is more elegant? 

- Look at the code for lab1 and write a function that prints the names of the registered workers periodically (every n seconds). 


Run the files with the following:

	$ cd excercises/session1
	$ go run duplicate.go words.txt
	$ go run sum.go

<!-- References -->
[1]: https://tour.golang.org/welcome/1 "Tour"
[2]: https://golang.org/doc/effective_go.html  "Effective Go"
[3]: https://golang.org/doc/effective_go.html#slices "Slices" 
[4]: https://golang.org/ref/spec#Channel_types "Channels"
[5]: http://golang.org/doc/code.html "How To Write Go Code"
