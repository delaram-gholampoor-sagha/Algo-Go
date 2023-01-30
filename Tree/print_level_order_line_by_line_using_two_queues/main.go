package main

import "fmt"

func (t *Tree) PrintLevelOrderLineByLine() {

	que1 := new(Queue)
	que2 := new(Queue)
	var temp *Node
	if t.root != nil {
		que1.Enqueue(t.root)
	}

	for que1.Length() != 0 || que2.Length() != 0 {

		for que1.Length() != 0 {
			temp2 := que1.Dequeue()
			temp = temp2.(*Node)
			fmt.Print(temp.value, " ")
			if temp.left != nil {
				que2.Enqueue(temp.left)
			}
			if temp.right != nil {
				que2.Enqueue(temp.right)
			}
		}

		fmt.Println(" ")

		for que2.Length() != 0 {
			temp2 := que2.Dequeue()
			temp = temp2.(*Node)
			fmt.Print(temp.value, " ")
			if temp.left != nil {
				que1.Enqueue(temp.left)
			}
			if temp.right != nil {
				que1.Enqueue(temp.right)
			}
		}

		fmt.Println(" ")

	}
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := LevelOrderBinaryTree(arr)
	t.PrintLevelOrderLineByLine()
}
