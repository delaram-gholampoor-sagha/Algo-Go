package main

func (t *Tree) TrimOutsidedataRange(min int, max int) {
	t.root = trimOutsidedataRange(t.root, min, max)
}

func trimOutsidedataRange(curr *Node, min int, max int) *Node {
	if curr == nil {
		return nil
	}
	curr.left = trimOutsidedataRange(curr.left, min, max)
	curr.right = trimOutsidedataRange(curr.right, min, max)
	if curr.value < min {
		return curr.right
	}
	if curr.value > max {
		return curr.left
	}
	return curr
}

/* Testing Code */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := CreateBinarySearchTree(arr)
	t.TrimOutsidedataRange(3, 9)
	t.PrintInOrder()
}
