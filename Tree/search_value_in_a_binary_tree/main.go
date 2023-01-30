package main
import "fmt"

func (t *Tree) SearchBT(value int) bool {
    return searchBT(t.root, value)
}

func searchBT(root *Node, value int) bool {
    var left, right bool
    if root == nil {
        return false
    }

    if root.value == value {
        return true
    }

    left = searchBT(root.left, value)
    if left {
        return true
    }

    right = searchBT(root.right, value)
    if right {
        return true
    }
    
    return false
}

/* Testing Code */
func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    t := LevelOrderBinaryTree(arr)
    fmt.Print(t.SearchBT(9))
}