package main

import (
	"fmt"
	"sort"
)

func main() {
	number1 := 234
	number2 := 143
	slice1 := make([]int, 3)
	slice2 := make([]int, 3)
	for number1 > 0 {
		slice1 = append(slice1, number1%10)
		number1 = number1 / 10
	}
	for number2 > 0 {
		slice2 = append(slice2, number1%10)
		number2 = number2 / 10
	}

	sort.Ints(slice1)
	sort.Ints(slice2)
	var flag = 0
	if len(slice1) == len(slice2) {
		for i := 0; i < len(slice1); i++ {
			if slice1[i] != slice2[i] {
				flag = 1
			}
		}
		if flag == 0 {
			fmt.Println("Number are anagram")
		} else {
			fmt.Println("Number are not anagram")
		}
	} else {
		fmt.Println("Not anagram number")
	}

}
