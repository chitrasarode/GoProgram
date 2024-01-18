package main

import (
	"fmt"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)
	go helloworld()

	go goodbye()

}

func helloworld() {
	fmt.Println("Hello World!")
}

func goodbye() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Good Bye!")
}
