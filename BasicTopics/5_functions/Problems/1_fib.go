/*
Fibonacci Series:

Create a function to generate the Fibonacci series up to a specified number of terms. The Fibonacci series is a sequence of numbers

	where each number is the sum of the two preceding ones, usually starting with 0 and 1.
*/package main

import "fmt"

/*Algorithm
F(n)=F(n−1)+F(n−2)

with initial conditions:
F(0)=0,F(1)=1
*/
func main() {
	var no int
	fmt.Println("-----Fibonacci series-----")
	fmt.Println("Enetr a number to get fibonacci series")
	fmt.Scan(&no)

	//without recursion
	f := fib(no)
	fmt.Println("Fibonacci numbers : ", f)

	//with recursion

	for i := 0; i < no; i++ {
		f1 := fib1(i)
		fmt.Printf("%d ", f1)
	}

}
func fib(no int) []int {
	f := make([]int, no)
	if no <= 1 {
		return f
	}
	f[0] = 0
	f[1] = 1

	for i := 2; i < no; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f
}
func fib1(no int) int {
	if no <= 1 {
		return no
	}

	return fib1(no-1) + fib1(no-2)
}
