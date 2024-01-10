package main

import (
	"fmt"
)

func main() {
	fmt.Println("Even array element")

	arr := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	even := []int{}
	fmt.Println("Array : ", arr)

	for _, v := range arr {
		if v%2 == 0 {
			even = append(even, v)
		}
	}
	fmt.Println("Even array element : ", even)
}
