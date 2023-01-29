package main
import "fmt"

func (t *Tree) PrintPreOrder() {
 fmt.Print("Tree : ")
 printPreOrder(t.root)
 fmt.Println()
}

func printPreOrder(n *Node) {
 if n == nil {
 return
 }
 fmt.Print(n.value, " ")
 printPreOrder(n.left)
 printPreOrder(n.right)
}

func main() {
    
    
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    t := LevelOrderBinaryTree(arr)
    t.PrintPreOrder()
}