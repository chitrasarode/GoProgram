package main

import (
	"fmt"
)

type Person struct {
	fname   string
	lname   string
	age     int
	Address Address
}
type Address struct {
	city string
	road string
	pin  int
}

func main() {
	fmt.Println("\nStructure Implementation")

	p1 := Person{
		fname: "Chitra",
		lname: "Sarode",
		age:   29,
		Address: Address{
			city: "Pune",
			road: "Ivy estate road",
			pin:  413201,
		},
	}

	p2 := Person{
		fname: "Rajashri",
		lname: "Chatrikar",
		age:   28,
		Address: Address{
			city: "Banglore",
			road: "Banglore highway",
			pin:  623145,
		},
	}

	printData(p1)
	printData(p2)
}
func printData(p Person) {
	fmt.Println("\nFirst Name : ", p.fname)
	fmt.Println("Last Name : ", p.lname)
	fmt.Println("Age : ", p.age)
	fmt.Println("Address : ", p.Address)
}
