package main

import "fmt"

func (t *Tree) PrintAllPath() {
	stk := new(Stack)
	printAllPath(t.root, stk)
}

func printAllPath(curr *Node, stk *Stack) {

	if curr == nil {
		return
	}

	stk.Push(curr.value)
	if curr.left == nil && curr.right == nil {
		stk.Print()
		fmt.Println()
		stk.Pop()
		return
	}

	printAllPath(curr.right, stk)
	printAllPath(curr.left, stk)

	stk.Pop()
}

/* Testing Code */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := LevelOrderBinaryTree(arr)
	t.PrintAllPath()
}
