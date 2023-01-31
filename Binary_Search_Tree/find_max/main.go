package main

import "fmt"

func (t *Tree) FindMaxNode() *Node {
	var node *Node = t.root
	if node == nil {
		fmt.Println("EmptyTreeException")
		return nil
	}

	for node.right != nil {
		node = node.right
	}
	return node
}

/* Testing Code */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := CreateBinarySearchTree(arr)
	maxNode := t.FindMaxNode()
	fmt.Println("Maximum :: ", maxNode.value)
}
