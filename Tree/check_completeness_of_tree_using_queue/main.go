package main

import "fmt"

func (t *Tree) IsCompleteTree() bool {
	return isCompleteTree(t.root)
}

func isCompleteTree(root *Node) bool {
	que := new(Queue)
	var temp *Node
	var noChild = false
	if root != nil {
		que.Enqueue(root)
	}
	for que.Length() != 0 {
		temp = que.Dequeue().(*Node)
		if temp.left != nil {
			if noChild == true {
				return false
			}
			que.Enqueue(temp.left)
		} else {
			noChild = true
		}
		if temp.right != nil {
			if noChild == true {
				return false
			}
			que.Enqueue(temp.right)
		} else {
			noChild = true
		}
	}
	return true
}

/* Testing Code */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := LevelOrderBinaryTree(arr)
	fmt.Println(t.IsCompleteTree())
}
