package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")
	go add(2, 3)
	go sub(10, 5)
	//duration := 100 * time.Millisecond
	//time.Sleep(duration)
	//time.Sleep(100)
	fmt.Println("End")
}
func add(a, b int) {
	Result := a + b
	fmt.Println("Addition : ", Result)
}
func sub(a, b int) {
	Result := a - b
	fmt.Println("Subtraction : ", Result)
}
