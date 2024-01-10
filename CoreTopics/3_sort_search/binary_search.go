package main

import "fmt"

// binarySearch performs a binary search on the given sorted slice.
// It returns the index of the target element if found, or -1 if not found.
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid // Target found, return its index
		} else if arr[mid] < target {
			low = mid + 1 // Target is in the right half
		} else {
			high = mid - 1 // Target is in the left half
		}
	}

	return -1 // Target not found
}

func main() {
	// Example usage
	numbers := []int{1, 2, 3, 4, 5, 7, 9}
	target := 7

	// Perform binary search
	index := binarySearch(numbers, target)

	// Display the result
	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the list\n", target)
	}
}
