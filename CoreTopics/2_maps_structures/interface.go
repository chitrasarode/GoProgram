package main

import (
	"fmt"
)

func main() {

	var value interface{}

	value = 42
	fmt.Println("Interface : ", value)
	fmt.Println("Enter the number : ")

	_, err := fmt.Scan(value)
	if err != nil {
		fmt.Println("Error : ", err)
	}
	fmt.Println("Number : ", value)
}
