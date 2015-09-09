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

	var intAry [3]int = [3]int{1, 2, 3}

	var myStruct Record
	myStruct.id = 999
	myStruct.name = "my record"

	// alternate initialization based on field definition positon
	var myStruct2 Record = Record{1000, "my other record"}

	fmt.Printf("%v %v %v %v %v %v %v %v\n", s, f, i, ptrToInt, intAry, i2, myStruct, myStruct2)
}
