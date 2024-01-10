package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slice Programs")
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//slice1 := slice[0:2]
	fmt.Println("Length of slice", len(slice))
	fmt.Printf("address of slice :%p\n", &slice)
	slice = append(slice, 10, 11)

	fmt.Printf("slice : %d\n", slice)
	fmt.Println("Length of slice", len(slice))
	fmt.Printf("address of slice : %p", &slice)
}
