package main

import "fmt"

func (t *Tree) TreeToListRec() *Node {
	return treeToListRec(t.root)
}

func treeToListRec(curr *Node) *Node {
	if curr == nil {
		return nil
	}
	var Head, Tail, tempHead *Node

	if curr.left == nil && curr.right == nil {
		curr.left = curr
		curr.right = curr
		return curr
	}

	if curr.left != nil {
		Head = treeToListRec(curr.left)
		Tail = Head.left
		curr.left = Tail
		Tail.right = curr
	} else {
		Head = curr
	}

	if curr.right != nil {
		tempHead = treeToListRec(curr.right)
		Tail = tempHead.left
		curr.right = tempHead
		tempHead.left = curr
	} else {
		Tail = curr
	}
	Head.left = Tail

	Tail.right = Head
	return Head
}

/* Testing Code */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := LevelOrderBinaryTree(arr)
	fmt.Print("Tree In Order Traversel :: ")
	t.PrintInOrder()
	fmt.Print("In Order Traverseled Linked List :: ")
	t3 := t.TreeToListRec()
	t3.PrintDLL()
}
