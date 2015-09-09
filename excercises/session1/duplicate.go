package main

import "fmt"
import "os"
import "io/ioutil"
import "strings"
import "path/filepath"

// Problem: given a string, find the words that contain duplicate instances.
// Define a word as a sequence of contiguous characters where each character
// is not a space or newline (/n or /r/n).
// Hint: the 'strings' package may be helpful
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
		fmt.Printf("usage: %v <filename>\n", filepath.Base(os.Args[0]))
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
