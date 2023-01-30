package main

import "fmt"

func (t *Tree) PrintDepthFirst() {

	stk := new(Stack)
	if t.root != nil {
		stk.Push(t.root)
	}

	fmt.Print("Depth First : ")

	for stk.Length() != 0 {

		temp := stk.Pop()
		fmt.Print(temp.value, " ")

		if temp.right != nil {
			stk.Push(temp.right)
		}

		if temp.left != nil {
			stk.Push(temp.left)
		}

	}
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := LevelOrderBinaryTree(arr)
	t.PrintDepthFirst()
}
