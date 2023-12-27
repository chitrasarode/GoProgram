package main

import (
	"fmt"
	"mainFolder/iLearn/iLearn/GoLang/BasicTopics/6_packages/Problems/calculator"
	"os"
)

func main() {
	var no1, no2 int
	var op string

	for {
		fmt.Println("1. +\n2. -\n3. *\n4. /\n5. Exit")
		fmt.Print("Enter operator : ")
		fmt.Scanln(&op)

		if op == "Exit" {
			os.Exit(1)
		} else if op != "+" && op != "-" && op != "*" && op != "/" {
			fmt.Print("Please enter correct operator :")
			fmt.Scanln(&op)
		}
		fmt.Print("Enter 1st number : ")
		fmt.Scanln(&no1)

		fmt.Print("Enter 2nd number : ")
		fmt.Scanln(&no2)

		switch op {
		case "+":
			calculator.Addition(no1, no2)
		case "-":
			calculator.Subtraction(no1, no2)
		case "*":
			calculator.Multiplication(no1, no2)
		case "/":
			calculator.Division(no1, no2)
		default:
			fmt.Print("Please Enter correct operator")

		}
	}

}
