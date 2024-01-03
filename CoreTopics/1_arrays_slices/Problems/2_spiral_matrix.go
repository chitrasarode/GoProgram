// The Spiral Matrix problem involves traversing a 2D matrix in a spiral order, starting from the top-left corner
// and moving in a clockwise direction. The task is to return all elements of the matrix in the order of their
// traversal. Here's a more detailed explanation:

// # Problem Statement:
// Given an m x n matrix, return all elements of the matrix in spiral order.

// # Example:
// Consider the following 3x3 matrix:

// 1 2 3
// 4 5 6
// 7 8 9

// The spiral order traversal of this matrix would be [1, 2, 3, 6, 9, 8, 7, 4, 5].

package main

import (
	"fmt"
)

func main() {
	var m, n int

	fmt.Println("-----Spiral Matrix Program-----")

	fmt.Println("Enter number of matrix rows")
	fmt.Scan(&m)

	fmt.Println("Enter number of matrix columns")
	fmt.Scan(&n)

	//Declare 2d matrix
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	var element int

	//To get matrix values
	fmt.Println("Enter matrix element : ")
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&element)
			matrix[i][j] = element
		}
	}

	//Print matrix values
	fmt.Println("Matrix : ")
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j], "\t")
		}
		fmt.Println()
	}
	var right int
	var bottom int

	//spiral order traversal of this matrix
	for i := 0; i < m; i++ {
		if right == n-1 {
			fmt.Print(matrix[i][right], ",")
			bottom = m - 1
			if bottom == i {
				for j := n - 2; j >= 0; j-- {
					fmt.Print(matrix[i][j], ",")
				}
			}
			continue
		}
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j], ",")
			right = n - 1
		}
	}
}
