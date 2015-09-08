package main

import "fmt"       // for Printf
import "os"        // to access the command line params
import "io/ioutil" // for ReadFile
import "strings"

// Problem: given a string, find the words that contain duplicate instances.
// Define a word as a sequence of characters that does not match
// the regular expression "[\r\n ]*".
func FindDuplicates(input string) []string {
	// your code here
	return []string{}
}

func GetFileContents(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %v <filename>\n", os.Args[0])
		os.Exit(1)
	}

	inputFile := os.Args[1]

	input := GetFileContents(inputFile)
	duplicates := FindDuplicates(input)

	fmt.Printf("input: %v\n", strings.Trim(input, "\r\n"))
	fmt.Printf("duplicates: ")
	for _, word := range duplicates {
		fmt.Printf("%v ", word)
	}
	fmt.Printf("\n\n")
}
