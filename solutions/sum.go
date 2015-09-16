package main

import "fmt" // for Printf
import "math/rand"
import "time"
import "runtime"
import "math"

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

func SumParallelWorker(a []int, result chan int) {
	result <- SimpleSum(a)
}

func SumParallel(a []int) int {
	blockSize := int(math.Ceil(float64(len(a)) / float64(numPartitions)))
	block := 0
	result := make(chan int)

	for block < numPartitions {
		var end int
		start := block * blockSize
		if block == numPartitions-1 {
			end = len(a)
		} else {
			end = block*blockSize + blockSize
		}
		go SumParallelWorker(a[start:end], result)
		block = block + 1
	}

	sum := 0
	cnt := 0
loop:
	for {
		select {
		case val := <-result:
			sum = sum + val
			cnt = cnt + 1
			if cnt == numPartitions {
				break loop
			}
		}
	}
	return sum
}

func initData(num int) []int {
	var result []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < num; i++ {
		result = append(result, r.Intn(Max))
	}
	fmt.Printf("array has %v items\n", num)
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
