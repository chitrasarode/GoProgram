package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	freq := make(map[string]int)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter line of string")
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input : ", err)
	}
	fmt.Println("line : ", line)
	word := strings.Fields(line)

	for _, v := range word {
		freq[v]++
	}
	for key, value := range freq {
		fmt.Printf("%s : %d\n", key, value)
	}
	// count := strings.Count(line, "chitra")
	// fmt.Println(count)

	// str := strings.Repeat(line, 2)//repeating string 2 times
	// fmt.Println(str)

	str1 := strings.Trim(line, " ")
	fmt.Println(str1)
}
