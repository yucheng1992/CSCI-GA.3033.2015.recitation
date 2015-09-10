package main

import "fmt"       // for Printf
import "os"        // to access the command line params
import "io/ioutil" // for ReadFile
import "strings"

// Problem: given a string, find the words that contain duplicate instances
// Note: A word will be defined as a sequence of characters that matches
// the regular expression "[A-Za-z0-9]*". The input should contain words
// separated by a space, carriage return, or newline.
func FindDuplicates(input string) []string {
	var duplicates []string // return value
	wordCount := make(map[string]int)
	words := strings.Split(input, " ")

	for _, word := range words {
		word = strings.Trim(word, "\r\n")
		if count, found := wordCount[word]; !found {
			wordCount[word] = 1
		} else {
			wordCount[word] = count + 1
		}
	}

	for word, count := range wordCount {
		if count > 1 {
			duplicates = append(duplicates, word)
		}
	}

	return duplicates
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
