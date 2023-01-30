package main

import (
	"fmt"
	"math"
)

func (t *Tree) FindMaxBT() int {
	return findMaxBT(t.root)
}

func findMaxBT(curr *Node) int {
	if curr == nil {
		return math.MinInt32
	}
	max := curr.value
	left := findMaxBT(curr.left)
	right := findMaxBT(curr.right)

	if left > max {
		max = left
	}
	if right > max {
		max = right
	}

	return max
}

/* Testing Code */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := LevelOrderBinaryTree(arr)
	fmt.Println(t.FindMaxBT())
}
