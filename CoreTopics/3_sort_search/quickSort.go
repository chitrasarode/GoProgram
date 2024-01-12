package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	var less, equal, greater []int
	for _, num := range arr {
		if num == pivot {
			equal = append(equal, num)
		} else if num < pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}

	}

	less = quickSort(less)
	greater = quickSort(greater)

	return append(append(less, equal...), greater...)
}

func main() {
	arr := []int{12, 13, 9, 3, 21, 16, 7, 25, 6}
	fmt.Println("Original array:", arr)

	sortedArr := quickSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
