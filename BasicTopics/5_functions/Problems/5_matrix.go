/*
Matrix Multiplication:

Implement a function that performs matrix multiplication. Given two matrices as input, return the resulting

	matrix. Make sure to check the compatibility of the matrices for multiplication.
*/
package main

import (
	"fmt"
)

var arr1 = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
var arr2 = [3][3]int{{7, 8, 9}, {6, 5, 4}, {3, 2, 1}}

func main() {

	fmt.Println("-----Matrix Multiplication Program-----")

	var arr3 = [3][3]int{}

	fmt.Println("1st Matrix : ")
	printMatrix(arr1)

	fmt.Println("2nd Matrix : ")
	printMatrix(arr2)

	fmt.Println("Matrix Multiplication : ")
	matrixMultplication(arr3)

}

func printMatrix(arr [3][3]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d\t", arr[i][j])
		}
		fmt.Printf("\n")
	}

}
func matrixMultplication(arr [3][3]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				arr[i][j] += arr1[i][k] * arr2[k][j]
			}
			fmt.Printf("%d\t", arr[i][j])
		}
		fmt.Printf("\n")
	}
}
