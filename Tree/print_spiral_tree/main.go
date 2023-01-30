package main

import "fmt"

func (t *Tree) PrintSpiralTree() {
	stk1 := new(Stack)
	stk2 := new(Stack)
	var temp *Node
	if t.root != nil {
		stk1.Push(t.root)
	}
	for stk1.Length() != 0 || stk2.Length() != 0 {
		for stk1.Length() != 0 {
			temp = stk1.Pop()
			fmt.Print(temp.value, " ")
			if temp.left != nil {
				stk2.Push(temp.left)
			}
			if temp.right != nil {
				stk2.Push(temp.right)
			}
		}
		fmt.Println(" ")
		for stk2.Length() != 0 {
			temp = stk2.Pop()
			fmt.Print(temp.value, " ")
			if temp.right != nil {
				stk1.Push(temp.right)
			}
			if temp.left != nil {
				stk1.Push(temp.left)
			}

		}
		fmt.Println(" ")
	}
}

/* Testing Code */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := LevelOrderBinaryTree(arr)
	t.PrintSpiralTree()
}
