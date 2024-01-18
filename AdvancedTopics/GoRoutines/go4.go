package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	go helloWorld()
	time.Sleep(100 * time.Millisecond)
}

func helloWorld() {
	fmt.Println("\nHello World")
	fmt.Println("Hello World1")
	fmt.Println("Hello World2")
}
