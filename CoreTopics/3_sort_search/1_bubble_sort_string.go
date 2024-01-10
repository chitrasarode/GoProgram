package main

import (
	"fmt"
)

func main() {
	fmt.Println("------Bubble Sort-----")
	arr := []string{"c", "a", "A", "D"}
	fmt.Printf("Array : %T\n", arr)

	arr1 := bubbleSort(arr)
	fmt.Println("Ascending array : ", arr1)
}
func bubbleSort(arr []string) []string {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
