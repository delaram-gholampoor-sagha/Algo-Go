package main

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

func (t *Tree) NumNodes() int {
	return numNodes(t.root)
}

func numNodes(curr *Node) int {
	if curr == nil {
		return 0
	}
	return (1 + numNodes(curr.left) + numNodes(curr.right))
}

func (t *Tree) IsCompleteTree() bool {
	count := t.NumNodes()
	return isCompleteTreeUtil(t.root, 0, count)
}

func isCompleteTreeUtil(curr *Node, index int, count int) bool {
	if curr == nil {
		return true
	}
	if index > count {
		return false
	}
	return isCompleteTreeUtil(curr.left, index*2+1, count) &&
		isCompleteTreeUtil(curr.right, index*2+2, count)
}
