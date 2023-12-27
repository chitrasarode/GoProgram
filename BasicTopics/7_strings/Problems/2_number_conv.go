/*
Write a program that takes numeric input as a string, converts it to a numeric type using the strconv package,
performs a mathematical operation, and then converts the result back to a string for display.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("----String Conversion----")
	var str1 = "12"
	var str2 = "24"

	numStr1, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Println("Error in conversion", err)
	}

	numStr2, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Println("Error in conversion", err)
	}

	Result := numStr1 + numStr2
	ResultStr := strconv.Itoa(Result)
	fmt.Println("Result string : ", ResultStr)

}
