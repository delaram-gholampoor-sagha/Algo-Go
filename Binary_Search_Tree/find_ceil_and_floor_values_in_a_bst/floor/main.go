package main

import (
	"fmt"
	"math"
)

func (t *Tree) FloorBST(val int) int {
	curr := t.root
	floor := math.MaxInt32
	for curr != nil {
		if curr.value >= val {
			curr = curr.left
		} else {
			floor = curr.value
			curr = curr.right
		}
	}
	if floor == math.MaxInt32 {
		return -1
	}
	return floor
}

/* Testing Code */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	t := CreateBinarySearchTree(arr)
	fmt.Println(t.FloorBST(8))
}
