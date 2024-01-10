package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 24, 7, 12, 16, 8}
	fmt.Println("Slice : ", slice)

	var greatest, currentGreatest int

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			currentGreatest = slice[i]
		}
		if currentGreatest > greatest {
			greatest = currentGreatest
		}
	}
	fmt.Println("Greatest : ", greatest)
}
