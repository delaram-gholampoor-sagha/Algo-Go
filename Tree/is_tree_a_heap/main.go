package main

import "fmt"

func (t *Tree) IsHeap() bool {
	parentValue := -99999999
	return t.IsCompleteTree() && isHeapUtil(t.root, parentValue)
}

func isHeapUtil(curr *Node, parentValue int) bool {
	if curr == nil {
		return true
	}
	if curr.value < parentValue {
		return false
	}
	return isHeapUtil(curr.left, curr.value) && isHeapUtil(curr.right,
		curr.value)
}

/* Testing Code */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := LevelOrderBinaryTree(arr)
	fmt.Println(t.IsHeap())
}
