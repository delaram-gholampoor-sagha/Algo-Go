package main

import "fmt"

func (t *Tree) FindMinNode() *Node {
	var node *Node = t.root
	if node == nil {
		fmt.Println("EmptyTreeException")
		return nil
	}

	for node.left != nil {
		node = node.left
	}
	return node
}

/* Testing Code */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := CreateBinarySearchTree(arr)
	minNode := t.FindMinNode()
	println("Minimum in BST :: ", minNode.value)
}
