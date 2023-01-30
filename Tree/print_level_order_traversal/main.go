package main

import "fmt"

func (t *Tree) PrintBreadthFirst() {
	que := new(Queue)
	var temp *Node
	if t.root != nil {
		que.Enqueue(t.root)
	}
	fmt.Print("Breadth First : ")
	for que.Length() != 0 {
		temp2 := que.Dequeue()
		temp = temp2.(*Node)
		fmt.Print(temp.value, " ")
		if temp.left != nil {
			que.Enqueue(temp.left)
		}
		if temp.right != nil {
			que.Enqueue(temp.right)
		}
	}
	fmt.Println()
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := LevelOrderBinaryTree(arr)
	t.PrintBreadthFirst()
}
