package main
import "fmt"

func (t *Tree) PrintPostOrder() {
 fmt.Print("Post Order : ")
 printPostOrder(t.root)
 fmt.Println()
}

func printPostOrder(n *Node) {
 if n == nil {
 return
 }
 printPostOrder(n.left)
 printPostOrder(n.right)
 fmt.Print(n.value, " ")
}

func main() {
    
    
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    t := LevelOrderBinaryTree(arr)
    t.PrintPostOrder()
}