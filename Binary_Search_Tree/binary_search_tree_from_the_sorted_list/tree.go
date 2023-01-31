package main
import "fmt"

type Node struct {
    value int
    left, right *Node
}

type Tree struct {
 root *Node
}

func (t *Tree) PrintInOrder() {
    fmt.Print("In Order : ")
    printInOrder(t.root)
    fmt.Println()
}

func printInOrder(n *Node) {
    if n == nil {
    return
    }
    printInOrder(n.left)
    fmt.Print(n.value, " ")
    printInOrder(n.right)
}