package main

import "fmt"

func (t *Tree) TreeDepth() int {
	return treeDepth(t.root)
}

func treeDepth(root *Node) int {

	if root == nil {
		return -1
	}
	lDepth := treeDepth(root.left)
	rDepth := treeDepth(root.right)

	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}

/* Testing Code */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := LevelOrderBinaryTree(arr)
	fmt.Println(t.TreeDepth())
}
