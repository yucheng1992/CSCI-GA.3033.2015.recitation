package main

import "fmt"

func main() {
	var intAry = [4]int{1, 2, 3, 4}

	// creates a new slice; does not allocate a new backing array
	var intSlice = intAry[1:3] // contians elements [2,3]

	// changes to the slice are reflected in the array and vice-versa
	intSlice[1] = -1
	intAry[1] = -2

	// creates a new slice while allocating storage for a backing array
	var intSlice2 = make([]int, 4, 4) // contains [0 0 0 0]
	intSlice2[0] = -3
	intSlice2[1] = -2
	intSlice2[2] = -1

	// output: [1 -2 -1 4] [-2 -1] [-3 -2 -1 0]
	fmt.Printf("elements: %v %v %v\n", intAry, intSlice, intSlice2)

	// get the length of a slice (or array) with the 'len' builtin
	fmt.Printf("lengths: %v %v %v\n", len(intAry), len(intSlice), len(intSlice2))
}
