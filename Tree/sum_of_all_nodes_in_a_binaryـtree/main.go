package main

import "fmt"

func (t *Tree) SumAllBT() int {
	return sumAllBT(t.root)
}

func sumAllBT(curr *Node) int {
	var sum, leftSum, rightSum int

	if curr == nil {
		return 0
	}

	rightSum = sumAllBT(curr.right)
	leftSum = sumAllBT(curr.left)
	sum = rightSum + leftSum + curr.value
	return sum
}

/* Testing Code */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := LevelOrderBinaryTree(arr)
	fmt.Print(t.SumAllBT())
}
