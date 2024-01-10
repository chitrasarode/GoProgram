package main

import (
	"fmt"
)

func main() {
	fmt.Println("Basic map programs")

	var a map[string]string         //nil map
	var b = make(map[string]string) //empty map

	fmt.Println(a == nil)
	fmt.Println(b == nil)
	fmt.Println(len(a) == 0)
	fmt.Println(len(b) == 0)

}
