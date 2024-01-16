package main

import (
	"fmt"
)

type node struct {
	data   int
	lchild *node
	rchild *node
}
type tree struct {
	root *node
}

func main() {

	fmt.Println("Binary Search Tree Program")
	t := tree{}
	t.insert(12)
	t.insert(23)
	t.insert(40)
	t.insert(25)
	t.insert(35)
	t.insert(7)
	t.insert(15)
	fmt.Print("Pre Order Traversal : ")
	preOrder(t.root)
	fmt.Print("\nIn Order Traversal : ")
	inOrder(t.root)
	fmt.Print("\nPost Order Traversal : ")
	postOrder(t.root)

}

func (t *tree) insert(data int) {
	newNode := &node{
		data:   data,
		lchild: nil,
		rchild: nil,
	}
	if t.root == nil {
		t.root = newNode
	} else {
		t.insertNode(t.root, newNode)
	}
}

func (t *tree) insertNode(root, newNode *node) {
	if newNode.data < root.data {
		if root.lchild == nil {
			root.lchild = newNode
		} else {
			t.insertNode(root.lchild, newNode)

		}
	} else {
		if root.rchild == nil {
			root.rchild = newNode
		} else {
			t.insertNode(root.rchild, newNode)
		}
	}
}

func preOrder(node *node) {
	if node != nil {
		fmt.Print(node.data, " ")
		preOrder(node.lchild)

		preOrder(node.rchild)
	}
}
func inOrder(node *node) {
	if node != nil {
		inOrder(node.lchild)
		fmt.Print(node.data, " ")
		inOrder(node.rchild)
	}
}

func postOrder(node *node) {
	if node != nil {
		postOrder(node.lchild)
		postOrder(node.rchild)
		fmt.Print(node.data, " ")
	}
}
