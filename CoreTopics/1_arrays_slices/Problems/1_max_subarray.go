// Problem Statement:
// Given an array of integers, find the contiguous subarray (containing at least one number) with the largest sum and return the sum.

// Example:
// Input: [-2, 1, -3, 4, -1, 2, 1, -5, 4]

// Output: subarray [4, -1, 2, 1] maxSum 6

// Explanation: The contiguous subarray [4, -1, 2, 1] has the largest sum of 6.

package main

import (
	"fmt"
)

func main() {
	// var arr []int
	// arr1 := [7]int{1, 2, 3, 4, 5, -1, -2}
	// fmt.Printf("\nType of slice = %T\ntype of array = %T", arr, arr1)
	// fmt.Printf("\nlength of slice = %d\nlength of array = %d\n", len(arr), arr1)
	// mySlice := arr1[1:4]
	// fmt.Println("Length of mySlice = ", len(mySlice), "\nCapacity of my slice = ", cap(mySlice))

	arr := [9]int{-2, 1, -3, 4, -1, 2, 1, 5, 9}
	maxSum := arr[0]
	sum := arr[0]
	fmt.Println("Array is : ", arr)
	subArray := []int{}
	// for _, value := range arr {
	// 	sum = sum + value
	// 	if {

	// 	}
	// }
	//-2, 1, -3, 4, -1, 2, 1, -5, 8

	var start int
	var tempStart int
	var end int
	for i := 1; i < len(arr); i++ {

		sum = arr[i] + sum
		if arr[i] > sum {
			sum = arr[i]
			tempStart = i
		} else {
			sum = sum + arr[i]
		}
		if sum > maxSum {
			maxSum = sum
			start = tempStart
			end = i
		}

	}
	subArray = arr[start : end+1]
	fmt.Println(maxSum)
	fmt.Println(subArray)
}
