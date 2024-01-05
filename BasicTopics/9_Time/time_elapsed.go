package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	fmt.Println("-----Fibonacci Series Program-----")
	fmt.Println("Start Time = ", startTime)
	var num int

	fmt.Println("Enter number to print the fibonacci series")
	fmt.Scan(&num)
	series := fib(num)
	fmt.Println(series)

	endTime := time.Now()
	fmt.Println("End Time = ", endTime)

	elapsedtime := endTime.Sub(startTime)
	fmt.Println("Elapsed time is = ", elapsedtime)

}

func fib(num int) []int {
	series := make([]int, num)
	series[0] = 0
	series[1] = 1
	for i := 2; i < num; i++ {
		series[i] = series[i-2] + series[i-1]
	}
	return series
}
