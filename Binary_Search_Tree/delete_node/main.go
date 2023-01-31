package main

import "fmt"

func (t *Tree) DeleteNode(value int) {
	t.root = DeleteNode(t.root, value)
}

func DeleteNode(node *Node, value int) *Node {
	var temp *Node = nil
	if node != nil {
		if node.value == value {
			if node.left == nil && node.right == nil {
				return nil
			}
			if node.left == nil {
				temp = node.right
				return temp
			}
			if node.right == nil {
				temp = node.left
				return temp
			}
			maxNode := FindMax(node.left)
			maxValue := maxNode.value
			node.value = maxValue
			node.left = DeleteNode(node.left, maxValue)
		} else {
			if node.value > value {
				node.left = DeleteNode(node.left, value)
			} else {
				node.right = DeleteNode(node.right, value)
			}
		}
	}
	return node
}

/* Testing Code */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := CreateBinarySearchTree(arr)
	t.PrintInOrder()
	t.DeleteNode(8)
	fmt.Println("After delete operation.")
	t.PrintInOrder()
}
