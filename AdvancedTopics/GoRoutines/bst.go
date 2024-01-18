package main

import (
	"fmt"
)

type node struct {
	data   int
	lChild *node
	rChild *node
}

type tree struct {
	root *node
}

func main() {
	fmt.Println("Binary Search Tree")

	t := tree{}
	t.insert(11)
	t.insert(21)
	t.insert(6)
	t.insert(16)
	t.insert(15)
	t.insert(9)
	t.insert(5)
	t.insert(22)
	t.search(22)
	t.delete(11)

	fmt.Print("\nPre Order : ")
	preOrder(t.root)
	fmt.Print("\nIn Order : ")
	InOrder(t.root)
	fmt.Print("\nPost Order : ")
	postOrder(t.root)

}

func (t *tree) insert(data int) {
	newNode := &node{
		data:   data,
		lChild: nil,
		rChild: nil,
	}
	if t.root == nil {
		t.root = newNode
	} else {
		t.insertNode(t.root, newNode)
	}
}
func (t *tree) insertNode(root, newNode *node) {
	if newNode.data < root.data {
		if root.lChild == nil {
			root.lChild = newNode
		} else {
			t.insertNode(root.lChild, newNode)
		}

	} else {
		if root.rChild == nil {
			root.rChild = newNode
		} else {
			t.insertNode(root.rChild, newNode)
		}
	}
}

func preOrder(node *node) {
	if node != nil {
		fmt.Print(node.data, " ")
		preOrder(node.lChild)
		preOrder(node.rChild)
	}
}

func postOrder(node *node) {
	if node != nil {
		postOrder(node.lChild)
		postOrder(node.rChild)
		fmt.Print(node.data, " ")
	}
}
func InOrder(node *node) {

	if node != nil {
		InOrder(node.lChild)
		fmt.Print(node.data, " ")
		InOrder(node.rChild)
	}
}
func (t *tree) search(key int) {
	if t.root == nil {
		fmt.Println("Tree is empty")
		return
	}
	current := t.root
	for current != nil {
		if key == current.data {
			fmt.Println("Element found")
			return
		} else if key < current.data {
			current = current.lChild
		} else {
			current = current.rChild
		}
	}
	fmt.Println("not found")
}
func (t *tree) delete(key int) {
	if t.root == nil {
		fmt.Println("Tree is empty")
		return
	}
	current := t.root
	for current != nil {
		if key == current.data {
			fmt.Println("Element found")
			return
		} else if key < current.data {
			current = current.lChild
		} else {
			current = current.rChild
		}
	}
	fmt.Println("not found")

}
