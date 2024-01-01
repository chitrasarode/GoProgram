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
		fmt.Println("1.Addition\n2.Subtraction\n3.Multiplication\n4.Division")
		fmt.Println("Enter your choice")
		fmt.Scan(&ch)

		if ch != 1 && ch != 2 && ch != 3 && ch != 4 {
			fmt.Println("Enter correct choice")
			continue
		}

		fmt.Println("Enter first number")
		fmt.Scan(&num1)

		fmt.Println("Enter second number")
		fmt.Scan(&num2)

		switch ch {
		case 1:
			fmt.Println("Addition of two numbers : ", calci.Addition(num1, num2))
		case 2:
			calci.Subtraction(num1, num2)
		case 3:
			calci.Multiplication(num1, num2)
		case 4:
			calci.Division(num1, num2)
		default:
			os.Exit(1)
		}
	}
}
