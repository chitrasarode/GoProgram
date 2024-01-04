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

	spiralMatrix := spiral_Matrix(matrix)
	fmt.Println("SpiralSlice : ", spiralMatrix)
}

func spiral_Matrix(matrix [][]int) []int {
	var right int
	var bottom int
	var top int
	var m = len(matrix)
	var n = len(matrix[0])
	spiralSlice := []int{}
	//var top int
	//spiral order traversal of this matrix
	for i := 0; i < m; i++ {
		if right == n-1 {
			spiralSlice = append(spiralSlice, matrix[i][right])
			fmt.Print(matrix[i][right], ",")
			bottom = m - 1
			if bottom == i {
				for j := n - 2; j >= 0; j-- {
					spiralSlice = append(spiralSlice, matrix[i][j])
					fmt.Print(matrix[i][j], ",")
					top = 0
				}
				if top == 0 {
					for i := n - 2; i > 0; i-- {
						spiralSlice = append(spiralSlice, matrix[i][top])
						fmt.Print(matrix[i][top], ",")
					}
				}
			}
			continue
		}
		for j := 0; j < n; j++ {
			spiralSlice = append(spiralSlice, matrix[i][j])
			fmt.Print(matrix[i][j], ",")
			right = n - 1
		}
	}
	// innerMatrix := make([][]int, len(matrix)-2)

	// fmt.Println(innerMatrix)
	// if len(innerMatrix) < 2 {
	// 	return spiralSlice
	// }
	return spiralSlice
}
