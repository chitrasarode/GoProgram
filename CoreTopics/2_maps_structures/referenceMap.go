package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps are references")

	var a = map[string]int{"chitra": 25, "varsha": 20, "nutan": 28, "rajashri": 24}
	// b := a
	fmt.Println("Before calling function map a is: ", a)
	// fmt.Println("map b: ", b)

	// b["varsha"] = 30
	// fmt.Println("map a: ", a)
	// fmt.Println("map b: ", b)

	referenceMap(a)
	fmt.Println("After calling function map a is: ", a)
}

func referenceMap(a map[string]int) {
	b := a
	fmt.Println("Map inside a function : ", b)
	b["chitra"] = 20
}
