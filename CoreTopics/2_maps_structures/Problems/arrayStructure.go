package main

import (
	"fmt"
)

type student struct {
	roll_no       int
	name, Address string
	marks         []int
}

func main() {

	fmt.Println("Array Structure")
	var s = []student{}
	var num int
	fmt.Println("Enter number of student ")
	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		var currentStudent student
		fmt.Println("Enter student roll no ")
		fmt.Scan(&currentStudent.roll_no)
		fmt.Println("Enter student name ")
		fmt.Scan(&currentStudent.name)
		fmt.Println("Enter student address ")
		fmt.Scan(&currentStudent.Address)
		fmt.Println("Enter student marks for 4 subject ")
		currentStudent.marks = make([]int, 4)
		for j := 1; j <= 4; j++ {
			fmt.Printf("Subject%d marks :", j)
			fmt.Scan(&currentStudent.marks[j-1])
		}
		s = append(s, currentStudent)
	}
	printStudentData(s)
}

func printStudentData(s []student) {
	for i, data := range s {
		fmt.Printf("\nStudent %d Information", i+1)
		fmt.Println("\n\tStudent Roll No : ", data.roll_no)
		fmt.Println("\tStudent Name : ", data.name)
		fmt.Println("\tStudent Address : ", data.Address)
		fmt.Println("\tStudent Marks : ")
		for i, v := range data.marks {
			fmt.Printf("\t\tSubject%d : %d\n", i+1, v)
		}
	}
}
