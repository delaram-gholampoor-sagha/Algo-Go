package main

import "fmt"

func (t *Tree) Free() {
	t.root = nil
}

/* Testing Code */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := LevelOrderBinaryTree(arr)
	t.Free()
	if t.root == nil {
		fmt.Println("The tree has been emptied.")
	} else {
		fmt.Println("The tree has not been emptied.")
	}
}
