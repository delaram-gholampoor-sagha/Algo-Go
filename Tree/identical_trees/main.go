package main
import "fmt"

func (t *Tree) IsEqual(t2 *Tree) bool {
    return isEqual(t.root, t2.root)
}

func isEqual(node1 *Node, node2 *Node) bool {
    if node1 == nil && node2 == nil {
        return true
    } else if node1 == nil || node2 == nil {
        return false
    } else {
        return ((node1.value == node2.value) &&
        isEqual(node1.left, node2.left) &&
        isEqual(node1.right, node2.right))
    }
}

/* Testing Code */
func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    t := LevelOrderBinaryTree(arr)
    t2 := LevelOrderBinaryTree(arr)
    fmt.Println(t.IsEqual(t2))
}