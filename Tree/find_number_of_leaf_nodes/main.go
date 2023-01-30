package main
import "fmt"

func (t *Tree) NumLeafNodes() int {
    return numLeafNodes(t.root)
}

func numLeafNodes(curr *Node) int {
    if curr == nil {
        return 0
    }
    if curr.left == nil && curr.right == nil {
        return 1
    }
    return (numLeafNodes(curr.right) + numLeafNodes(curr.left))
}

/* Testing Code */
func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    t := LevelOrderBinaryTree(arr)
    fmt.Println(t.NumLeafNodes())
}