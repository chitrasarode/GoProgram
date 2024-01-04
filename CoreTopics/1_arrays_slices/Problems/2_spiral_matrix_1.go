package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----Spiral Matrix-----")

	var m, n, element int

	fmt.Println("Enter number of rows")
	fmt.Scan(&m)

	fmt.Println("Enter number of columns")
	fmt.Scan(&n)

	//created 2d matrix
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	//To get matrix values
	fmt.Println("Enter Matrix element")
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&element)
			matrix[i][j] = element
		}
	}

	//To print matrix
	fmt.Println("Matrix : ")
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j], "\t")
		}
		fmt.Println()
	}

	spiralMatrix := spiral_Matrix1(matrix)
	fmt.Println("Spiral Matrix : ", spiralMatrix)

}
func spiral_Matrix1(matrix [][]int) []int {
	var rowStart, rowEnd, colStart, colEnd int
	var spiralMatrix = []int{}

	rowStart = 0
	colStart = 0
	rowEnd = len(matrix)
	colEnd = len(matrix[0])

	for rowStart < rowEnd && colStart < colEnd {
		//To Print first row
		for i := colStart; i < colEnd; i++ {
			spiralMatrix = append(spiralMatrix, matrix[rowStart][i])
		}
		rowStart++
		for i := rowStart; i < rowEnd; i++ {
			spiralMatrix = append(spiralMatrix, matrix[i][colEnd-1])
		}
		colEnd--
		if rowStart < rowEnd {
			for i := colEnd - 1; i >= colStart; i-- {
				spiralMatrix = append(spiralMatrix, matrix[rowEnd-1][i])
			}
			rowEnd--
		}
		if colStart < colEnd {
			for i := rowEnd - 1; i >= rowStart; i-- {
				spiralMatrix = append(spiralMatrix, matrix[i][colStart])
			}
			colStart++
		}
	}
	return spiralMatrix

}
