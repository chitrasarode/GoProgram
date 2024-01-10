package main

import (
	"fmt"
)

func main() {
	fmt.Println("Delete the item in array")

	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Array :", arr)
	arr1 := []int{}

	delete := 8
	flag := 0
	var index int
	for i := 0; i < len(arr); i++ {
		if arr[i] == delete {
			flag = 1
			index = i
			goto out
		}
	}
out:
	if flag == 1 {
		arr1 = append(arr[0:index], arr[index+1:]...)
		fmt.Println("Array : ", arr1)
	} else {
		fmt.Println("Element not found")
	}
}
