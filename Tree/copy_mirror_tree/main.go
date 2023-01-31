package main

import "fmt"

func (t *Tree) CopyMirrorTree() *Tree {
	tree := new(Tree)
	tree.root = copyMirrorTree(t.root)
	return tree
}

func copyMirrorTree(curr *Node) *Node {
	var temp *Node
	if curr != nil {
		temp = new(Node)
		temp.value = curr.value
		temp.right = copyMirrorTree(curr.left)
		temp.left = copyMirrorTree(curr.right)
		return temp
	}
	return nil
}

/* Testing Code */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := LevelOrderBinaryTree(arr)
	t2 := t.CopyMirrorTree()
	fmt.Println("Copied tree is a mirror tree :: ", t.IsMirror(t2))
	fmt.Println("Printing original tree ::")
	t.PrintLevelOrderLineByLine()
	fmt.Println("Printing mirror tree ::")
	t2.PrintLevelOrderLineByLine()
}
