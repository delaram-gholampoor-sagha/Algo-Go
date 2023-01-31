package main

import "fmt"

func (root *Node) PrintDLL() {
	if root == nil {
		return
	}
	curr := root
	tail := curr.left

	for curr != tail {
		fmt.Print(curr.value, " ")
		curr = curr.right
	}
	fmt.Print(curr.value)
}
