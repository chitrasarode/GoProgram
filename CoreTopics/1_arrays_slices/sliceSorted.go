package main

import (
	"fmt"
)

func main() {
	fmt.Println("Check array is sorted or not")
	var flag int
	//arr := [6]int{1, 4, 3, 2, 5, 6}
	arr := [6]int{7, 6, 5, 4, 3, 2} //for descending sorted
	if len(arr) < 2 {
		fmt.Println("Too short array to sort")
	}
	if arr[0] > arr[1] {
		goto out
	} else {
		fmt.Println("Array is not sorted")
		return
	}
out:
	for i := 2; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			flag = 1
		} else {
			flag = 0
			break
		}
	}
	if flag == 1 {
		fmt.Println("Array is sorted")
	} else {
		fmt.Println("Array is not sorted")
	}
	//for ascending sorted
	// 	//arr := [6]int{1, 4, 3, 2, 5, 6}
	// 	arr := [6]int{7, 6, 5, 4, 3, 2}
	// 	if len(arr) < 2 {
	// 		fmt.Println("Too short array to sort")
	// 	}
	// 	if arr[0] < arr[1] {
	// 		goto out
	// 	} else {
	// 		fmt.Println("Array is not sorted")
	// 		return
	// 	}
	// out:
	// 	for i := 2; i < len(arr); i++ {
	// 		if arr[i-1] < arr[i] {
	// 			flag = 1
	// 		} else {
	// 			flag = 0
	// 			break
	// 		}
	// 	}
	// 	if flag == 1 {
	// 		fmt.Println("Array is sorted")
	// 	} else {
	// 		fmt.Println("Array is not sorted")
	// 	}
}
