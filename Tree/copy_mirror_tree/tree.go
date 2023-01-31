package main

import "fmt"

type Node struct {
	value       int
	left, right *Node
}

type Tree struct {
	root *Node
}

func LevelOrderBinaryTree(arr []int) *Tree {
	tree := new(Tree)
	tree.root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}

func levelOrderBinaryTree(arr []int, start int, size int) *Node {
	curr := &Node{arr[start], nil, nil}
	left := 2*start + 1
	right := 2*start + 2
	if left < size {
		curr.left = levelOrderBinaryTree(arr, left, size)
	}
	if right < size {
		curr.right = levelOrderBinaryTree(arr, right, size)
	}
	return curr
}

func (t *Tree) PrintLevelOrderLineByLine() {
	que := new(Queue)
	var temp *Node
	var count int

	if t.root != nil {
		que.Enqueue(t.root)
	}
	for que.Length() != 0 {
		count = que.Length()
		for count > 0 {
			temp2 := que.Dequeue()
			temp = temp2.(*Node)
			fmt.Print(temp.value, " ")
			if temp.left != nil {
				que.Enqueue(temp.left)
			}
			if temp.right != nil {
				que.Enqueue(temp.right)
			}
			count -= 1
		}
		fmt.Println(" ")
	}
}

func (t *Tree) IsMirror(t2 *Tree) bool {
	return isMirror(t.root, t2.root)
}

func isMirror(node1 *Node, node2 *Node) bool {

	if node1 == nil && node2 == nil {
		return true
	} else if node1 == nil || node2 == nil {
		return false
	} else {
		return ((node1.value == node2.value) &&
			isMirror(node1.left, node2.right) &&
			isMirror(node1.right, node2.left))
	}
}
