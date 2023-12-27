package calculator

import "fmt"

func Addition(no1, no2 int) {
	addition := no1 + no2
	fmt.Println("Addition of two numbers : ", addition)
}
func Subtraction(no1, no2 int) {
	subtraction := no1 - no2
	fmt.Println("Subtraction of two numbers : ", subtraction)
}
func Multiplication(no1, no2 int) {
	multiplication := no1 * no2
	fmt.Println("Multiplication of two numbers : ", multiplication)
}
func Division(no1, no2 int) {
	if no2 == 0 {
		fmt.Println("division by zero error")
		return
	}
	division := no1 / no2
	fmt.Println("Division : ", division)
}
