package main

import (
	"fmt"
	"os"
)

type Stack struct {
	items []int
}

func main() {
	fmt.Println("Stack Implementation")

	stack := new(Stack)
	var ch int

	for {
		fmt.Println("\n1.Display Stack\n2.Insert element in stack\n3.Remove element from stack\n4.Exit")
		fmt.Print("Enter your choice : ")
		fmt.Scan(&ch)

		switch ch {
		case 1:
			stack.Display()
		case 2:
			var item int
			fmt.Print("Enter element to add in stack : ")
			fmt.Scan(&item)
			stack.push(item)
		case 3:
			stack.pop()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("Enter correct choice")
		}
	}
}

func (s *Stack) push(item int) {
	s.items = append(s.items, item)
	fmt.Println("Element inserted...")
}

func (s *Stack) Display() {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return
	}
	for i := len(s.items) - 1; i >= 0; i-- {
		if i == len(s.items)-1 {
			fmt.Println(s.items[i], "<-Top")
			continue
		}
		fmt.Println(s.items[i])
	}
}

func (s *Stack) pop() {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	fmt.Printf("%d deleted successfuly", item)
}
