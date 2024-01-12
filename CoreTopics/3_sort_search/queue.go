package main

import (
	"fmt"
	"os"
)

type queue struct {
	items []int
}

func main() {
	fmt.Println("\nQueue Program")
	var q = new(queue)
	var ch int

	for {
		fmt.Println("\n\n1.Display Queue\n2.Enqueue\n3.Dequeue\n4.Exit")
		fmt.Print("Enter your choice : ")
		fmt.Scan(&ch)
		switch ch {
		case 1:
			q.display()
		case 2:
			var item int
			fmt.Print("Enter element to add in queue : ")
			fmt.Scan(&item)
			q.Enqueue(item)
		case 3:
			q.Dequeue()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("Enter correct choice")
		}
	}
}

func (q *queue) display() {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return
	}
	fmt.Print("Queue elements are : ")
	for _, v := range q.items {
		fmt.Print(v, ",")
	}
}
func (q *queue) Enqueue(item int) {
	q.items = append(q.items, item)
	fmt.Println("Element inserted..")
}
func (q *queue) Dequeue() {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return
	}
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	fmt.Printf("%d deleted successfully", item)
}
