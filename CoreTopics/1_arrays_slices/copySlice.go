package main

import (
	"fmt"
)

func main() {
	fmt.Println()

	slice1 := []int{1, 2, 3}
	fmt.Println("slice1 : ", slice1)

	slice2 := []int{4, 5, 6}
	fmt.Println("slice1 : ", slice2)

	slice2 = append(slice1, slice2...)
	fmt.Println("copy slice1 and slice2 into slice3 : ", slice2)
}
