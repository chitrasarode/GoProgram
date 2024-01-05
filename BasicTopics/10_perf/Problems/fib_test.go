package main

import "testing"

func BenchmarkFibonacci(b *testing.B) {
	// Run the function b.N times
	for i := 0; i < b.N; i++ {
		fibonacci(10) // Adjust the input as needed
	}
}
