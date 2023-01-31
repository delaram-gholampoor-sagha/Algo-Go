package main

import "fmt"

func IsBST(root *Node) bool {
	if root == nil {
		return true
	}
	if root.left != nil && FindMax(root.left).value > root.value {
		return false
	}
	if root.right != nil && FindMin(root.right).value <= root.value {
		return false
	}

	return (IsBST(root.left) && IsBST(root.right))
}

/* Testing Code */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := CreateBinarySearchTree(arr)
	fmt.Println(IsBST(t.root))
}
