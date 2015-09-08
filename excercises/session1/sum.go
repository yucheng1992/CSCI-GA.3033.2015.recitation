package main

import "fmt" // for Printf
import "time"
import "runtime"
import "math/rand"

const (
	Size = 10E6
	Max  = 10
)

func SumSingleThread(a []int) int {
	s := 0
	for _, v := range a {
		s = s + v
	}
	return s
}

// returns the result of the sum of all elements in a.
// This should process the slice in multiple goroutines
func SumMultithread(a []int) int {
	// your code here
	return 0
}

func initData(num int) []int {
	var result []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < num; i++ {
		result = append(result, r.Intn(Max))
	}

	return result
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	a := initData(Size)

	start := time.Now()
	sum := SumMultithread(a)
	elapsedParallel := time.Since(start)

	start = time.Now()
	check := SumSingleThread(a)
	elapsedSingle := time.Since(start)

	if sum != check {
		fmt.Printf("sum incorrect\nexpected %v; got %v\n", check, sum)
	} else {
		fmt.Printf("sum %v calculated correctly\n", sum)
		fmt.Printf("parallel took %v\n", elapsedParallel)
		fmt.Printf("single took %v\n", elapsedSingle)
	}
}
