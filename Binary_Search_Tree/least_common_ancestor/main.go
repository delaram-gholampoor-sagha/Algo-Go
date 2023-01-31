package main

import "fmt"

func (t *Tree) LcaBST(first int, second int) (int, bool) {
	return LcaBST(t.root, first, second)
}

func LcaBST(curr *Node, first int, second int) (int, bool) {

	if curr == nil {
		fmt.Println("NotFoundException")
		return 0, false
	}
	if curr.value > first && curr.value > second {
		return LcaBST(curr.left, first, second)
	}
	if curr.value < first && curr.value < second {
		return LcaBST(curr.right, first, second)
	}
	return curr.value, true
}

/* Testing Code */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := CreateBinarySearchTree(arr)
	fmt.Println(t.LcaBST(3, 4))
	fmt.Println(t.LcaBST(1, 4))
	fmt.Println(t.LcaBST(10, 4))
}
