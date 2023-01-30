package main
import "fmt"

func (t *Tree) NumNodes() int {
    return numNodes(t.root)
}

func numNodes(curr *Node) int {
    //returns 0 when trying to traverse children of a leaf node 
    if curr == nil {
        return 0
    }
    // returns the count of nodes
    return (1 + numNodes(curr.right) + numNodes(curr.left))
}

/* Testing Code */
func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    t := LevelOrderBinaryTree(arr)
    fmt.Print(t.NumNodes())
}