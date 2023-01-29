package main

type Node struct {
 value int
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