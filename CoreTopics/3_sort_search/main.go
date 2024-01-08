package main

import (
	//bubbleSort "command-line-argumentsC:\\Users\\Admin\\Documents\\GoLang\\Go_Programs\\GoProgram\\CoreTopics\\3_sort_search\\bubbleSort\\1_bubble_sort_int.go"
	"fmt"
	"math/rand"
	"sorting/bubbleSort"
	"time"
)

func main() {
	fmt.Println("------Bubble Sort-----")

	arr := make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	fmt.Println("Array : ", arr)

	arr1 := bubbleSort.MyBubbleSort(bubbleSort.Ascending, arr)
	fmt.Println("Ascending Array : ", arr1)

	arr2 := bubbleSort.MyBubbleSort(bubbleSort.Descending, arr)
	fmt.Println("Descending Array : ", arr2)
}
