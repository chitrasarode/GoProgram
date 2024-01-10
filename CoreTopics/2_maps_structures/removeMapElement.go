package main

import (
	"fmt"
)

func main() {
	fmt.Println("Remove map element")

	var person = make(map[string]string)
	person["chitra"] = "pune"
	person["rajashri"] = "mumbai"
	person["shalini"] = "ranchi"

	fmt.Println(person)
	fmt.Println("Length of map : ", len(person))

	person["nitin"] = "banglore" //add element to map
	fmt.Println(person)
	fmt.Println("Length of map : ", len(person))

	person["chitra"] = "chennai" //update element
	fmt.Println(person)
	fmt.Println("Length of map : ", len(person))

	delete(person, "shalini") //delete element
	fmt.Println(person)
	fmt.Println("Length of map : ", len(person))

	person["nishchit"] = "Delhi" //add element to map
	fmt.Println(person)
	fmt.Println("Length of map : ", len(person))

	val, exist := person["rajashri"] //To check for specific element in map
	if exist {
		fmt.Println("This element exist in map : ", val)
	} else {
		fmt.Println("This element does exist in map")

	}
}
