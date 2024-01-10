package main

import (
	"fmt"
)

func main() {
	fmt.Println("Reverse the array")

	arr := [5]int{10, 23, 43, 21, 36}
	fmt.Println("Array : ", arr)

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println("Reversed array : ", arr)
}
