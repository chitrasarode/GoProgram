package main

import (
	"fmt"
	"time"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// func BenchmarkFibonacci(b *testing.B) {
// 	// Run the function b.N times
// 	for i := 0; i < b.N; i++ {
// 		fibonacci(10) // Adjust the input as needed
// 	}
// }

func main() {

	start := time.Now()

	fibonacci(20)

	end := time.Now()

	fmt.Printf("Time taken %s \n", end.Sub(start))
}
