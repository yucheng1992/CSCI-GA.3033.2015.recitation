package main

import "fmt" // for Printf
import "time"
import "runtime"
import "math/rand"

const (
	numPartitions = 4
	Size          = 10E6
	Max           = 10
)

// single threaded sum
func SimpleSum(a []int) int {
	s := 0
	for _, v := range a {
		s = s + v
	}
	return s
}

func SumParallel(a []int) int {
	return 0
}

func initData(num int) []int {
	var result []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < num; i++ {
		result = append(result, r.Intn(Max))
	}
	//fmt.Printf("array has %v items\n", num)
	return result
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	a := initData(Size)

	start := time.Now()
	sum := SumParallel(a)
	elapsedParallel := time.Since(start)

	start = time.Now()
	check := SimpleSum(a)
	elapsedSingle := time.Since(start)

	if sum != check {
		fmt.Printf("expected %v; got %v\n", check, sum)
	} else {
		fmt.Printf("sum %v calculated correctly\n", sum)
		fmt.Printf("parallel took %v\n", elapsedParallel)
		fmt.Printf("single took %v\n", elapsedSingle)
	}
}
