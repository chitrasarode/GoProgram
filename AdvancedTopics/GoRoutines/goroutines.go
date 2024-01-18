package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(500 * time.Millisecond) // Simulate some work
		fmt.Printf("%d ", i)
	}
}

func printLetters() {
	for char := 'a'; char <= 'e'; char++ {
		time.Sleep(300 * time.Millisecond) // Simulate some work
		fmt.Printf("%c ", char)
	}
}

func main() {
	fmt.Println("Start")

	// Launch goroutines concurrently
	go printNumbers()
	go printLetters()

	//Sleep to allow goroutines to complete
	time.Sleep(3 * time.Second)

	fmt.Println("\nEnd")
}
