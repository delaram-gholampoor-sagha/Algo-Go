package main
import "fmt"

func (t *Tree) IsCompleteTree2() bool {
    count := t.NumNodes()
    return isCompleteTreeUtil(t.root, 0, count)
}

func isCompleteTreeUtil(curr *Node, index int, count int) bool {
    if (curr == nil){
        return true
    }
    if (index > count){
        return false
    }
    return isCompleteTreeUtil(curr.left, index * 2 + 1, count) &&
    isCompleteTreeUtil(curr.right, index * 2 + 2, count)
}

/* Testing Code */
func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    t := LevelOrderBinaryTree(arr)
    fmt.Println(t.IsCompleteTree2());
}