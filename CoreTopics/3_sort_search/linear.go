package main

import (
	"fmt"
)

/*Algorithm

Every element is considered as a potential match for the key and checked for the same.
If any element is found equal to the key, the search is successful and the index of that element is returned.
If no element is found equal to the key, the search yields “No match found”.

*/

func main() {
	var arr1 = [4]int{11, 24, 36, 50}
	var x int = 24
	var n int = len(arr1)
	var i int
	var result bool = false

	fmt.Println("Arrray :  ", arr1)
	for i = 0; i < n; i++ {

		if arr1[i] == x {
			result = true
			fmt.Println("Element found at index ", i)
			break
		}
	}
	if result == false {
		fmt.Println("Element not found")
	}

}
