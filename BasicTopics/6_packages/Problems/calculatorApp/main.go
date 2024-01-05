package main

import (
	//calci "command-line-argumentsC:\\Users\\Admin\\Documents\\GoLang\\Go_Programs\\GoProgram\\BasicTopics\\6_packages\\Problems\\calculatorApp\\Calculator\\calci.go"
	"calculatorApp/calci"
	"fmt"
	"os"
)

func main() {
	var num1, num2, ch int

	for {
		fmt.Println("*****Calculator App*****")
		fmt.Println("1.Addition\n2.Subtraction\n3.Multiplication\n4.Division\n5.Exit")
		fmt.Println("Enter your choice")
		fmt.Scan(&ch)

		if ch < 1 || ch > 5 {
			fmt.Println("Enter correct choice")
			continue
		}

		fmt.Println("Enter first number")
		fmt.Scan(&num1)

		fmt.Println("Enter second number")
		fmt.Scan(&num2)

		switch ch {
		case 1:
			result, err := calci.Addition(num1, num2)
			if err != nil {
				fmt.Println("Error : ", err)
			}
			fmt.Println("Addition of two numbers : ", result)
		// case 2:
		// 	result, err := calci.Subtraction(num1, num2)
		// 	if err != nil {
		// 		fmt.Println("Error : ", err)
		// 	}
		// 	fmt.Println("Subtraction of two numbers : ", result)

		// case 3:
		// 	result, err := calci.Multiplication(num1, num2)
		// 	if err != nil {
		// 		fmt.Println("Error : ", err)
		// 	}
		// 	fmt.Println("Subtraction of two numbers : ", result)

		case 4:
			//calci.Division(num1, num2)
		case 5:
			os.Exit(1)
		}
	}
}
