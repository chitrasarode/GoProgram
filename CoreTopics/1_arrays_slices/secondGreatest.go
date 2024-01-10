package main

import (
	"fmt"
)

func main() {
	fmt.Println("Second Greatest number")

	arr := [9]int{5, 6, 7, 1, 2, 3, 4}

	fmt.Println(arr)

	var greatest = arr[0]
	var secondGreatest int

	for i := 0; i < len(arr); i++ {
		if arr[i] > greatest {
			secondGreatest = greatest
			greatest = arr[i]
		} //else if arr[i] > secondGreatest {
		// 	secondGreatest = greatest
		// }
	}
	fmt.Println(secondGreatest, greatest)
}
