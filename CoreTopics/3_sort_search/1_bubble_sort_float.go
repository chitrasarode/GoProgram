package main

import (
	"fmt"
)

func main() {
	fmt.Println("------Bubble Sort-----")

	arr := []float32{1.1, 1.0, 2.3, 2.1, 2.2, 3.1, 2.6}
	fmt.Println("Array : ", arr)

	arr1 := myBubbleSort1(ascending, arr)
	fmt.Println("Ascending Array : ", arr1)

	arr2 := myBubbleSort1(descending, arr)
	fmt.Println("Descending Array : ", arr2)

}

func myBubbleSort1(f func([]float32, int) []float32, arr []float32) []float32 {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			f(arr, j)
		}
	}
	return arr
}

func ascending(arr []float32, x int) []float32 {
	if arr[x] > arr[x+1] {
		arr[x], arr[x+1] = arr[x+1], arr[x]
	}
	return arr
}

func descending(arr []float32, x int) []float32 {
	if arr[x] < arr[x+1] {
		arr[x], arr[x+1] = arr[x+1], arr[x]
	}
	return arr
}
