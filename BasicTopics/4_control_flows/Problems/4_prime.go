/*
Prime number check: Check if given number is prime or not
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var no int

	fmt.Println("-----Prime Number Program-----")

	fmt.Print("Enter number to check it is prime or not : ")
	fmt.Scanln(&no)

	isPrime := prime(no)

	if isPrime {
		fmt.Println("Number is prime")
	} else {
		fmt.Println("Number is not prime")
	}
}
func prime(no int) bool {
	no1 := math.Sqrt(float64(no))
	for i := 2; i <= int(no1); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
