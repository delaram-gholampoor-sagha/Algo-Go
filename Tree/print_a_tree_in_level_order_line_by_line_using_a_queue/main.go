package main
import "fmt"

func (t *Tree) PrintLevelOrderLineByLine2() {
    que := new(Queue)
    var temp *Node
    var count int

    if t.root != nil {
        que.Enqueue(t.root)
    }
    for que.Length() != 0 {
        count = que.Length()
        for count > 0 {
            temp2 := que.Dequeue()
            temp = temp2.(*Node)
            fmt.Print(temp.value, " ")
            if temp.left != nil {
                que.Enqueue(temp.left)
            }
            if temp.right != nil {
                que.Enqueue(temp.right)
            }
            count -= 1
        }
        fmt.Println(" ")
    }
}

/* Testing Code */
func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    t := LevelOrderBinaryTree(arr)
    t.PrintLevelOrderLineByLine2()
}