// Word Frequency:

// Given a paragraph, write a program to find the frequency of each word. Store the word frequencies in a map.
// Example:
// go
// Copy code
// paragraph := "This is a sample paragraph. This paragraph contains some sample words."
// // Output: {"This": 2, "is": 1, "a": 1, "sample": 2, "paragraph": 2, "conta

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("---Word Frequency in paragraph---")
	paragraph := "This is a sample paragraph. This paragraph contains some sample words."
	//fmt.Print(paragraph)

	freq := map[string]int{}
	paragraph = strings.ReplaceAll(paragraph, ".", "")
	word := strings.Split(paragraph, " ")

	for _, v := range word {
		freq[v]++
	}
	for value, key := range freq {
		fmt.Printf(" %s  :  %d\n", value, key)
	}
}
