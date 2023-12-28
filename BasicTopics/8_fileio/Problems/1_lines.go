/*
Write a program that reads a text file and counts the number of lines in it. Use bufio.Reader for efficient line-by-line reading.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var count int

	file, err := os.Open("File1.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("Errror in opening file: ", err)
	}

	reader := bufio.NewReader(file)
	for {
		_, err := reader.ReadString('\n')
		//fmt.Println(content)
		count++
		if err != nil {
			break
		}
	}
	fmt.Println("Total number of lines in File : ", count)
}
