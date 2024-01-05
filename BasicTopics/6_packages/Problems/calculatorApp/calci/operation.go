package calci

import "errors"

//"fmt"

func Addition(num1, num2 int) (int, error) {
	if num1 < 0 || num2 < 0 {
		return 0, errors.New("negative numbers are not allowed")
	}
	result := num1 + num2
	return result, nil
}

// func Subtraction(no1, no2 int) {
// 	subtraction := no1 - no2
// 	fmt.Println("Subtraction of two numbers : ", subtraction)
// }
// func Multiplication(no1, no2 int) {
// 	multiplication := no1 * no2
// 	fmt.Println("Multiplication of two numbers : ", multiplication)
// }
// func Division(no1, no2 int) {
// 	if no2 == 0 {
// 		fmt.Println("division by zero error")
// 		return
// 	}
// 	division := no1 / no2
// 	fmt.Println("Division : ", division)
// }
