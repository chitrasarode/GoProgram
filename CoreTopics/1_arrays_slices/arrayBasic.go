package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array basic Programs")

	arr1 := [5]int{}              //not initialized
	arr2 := [5]int{1, 2, 3}       //partially initialized
	arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized
	arr4 := [5]int{2: 20, 4: 40}  //initialized specific elements

	fmt.Println("Array1 : ", arr1)
	fmt.Println("Array1 : ", arr2)
	fmt.Println("Array1 : ", arr3)
	fmt.Println("Array1 : ", arr4)

}
