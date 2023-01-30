package main
import "fmt"

func (t *Tree) NumFullNodesBT() int {
    return numFullNodesBT(t.root)
}

func numFullNodesBT(curr *Node) int {
    var count int

    if curr == nil {
        return 0
    }

    count = numFullNodesBT(curr.right) + numFullNodesBT(curr.left)
    
    if curr.right != nil && curr.left != nil {
        count++
    }
    return count
}

/* Testing Code */
func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    t := LevelOrderBinaryTree(arr)
    fmt.Println(t.NumFullNodesBT())
}