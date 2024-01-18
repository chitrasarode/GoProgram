package main

import (
	"fmt"
	"time"
)

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		//c <- i // Send value to the channel
		fmt.Printf("%d \n", i)
	}

	//close(c)
}

func main() {
	//numbersChannel := make(chan int)

	go PrintNumbers()
	time.Sleep(500 * time.Millisecond)
	for i := 'a'; i <= 'e'; i++ {
		//time.Sleep(100 * time.Millisecond)
		fmt.Printf("%c \n", i)
		//c <- i // Send value to the channel
	}

	// Output may vary on each run due to concurrent execution
}
