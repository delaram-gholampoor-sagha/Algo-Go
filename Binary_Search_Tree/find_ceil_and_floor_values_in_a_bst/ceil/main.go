package main
import "fmt"
import "math"

func (t *Tree) CeilBST(val int) int {
    curr := t.root
    ceil := math.MinInt32
    
    for curr != nil {
        if curr.value <= val {
        curr = curr.right
        } else {
        ceil = curr.value
        curr = curr.left
        }
    }
    if ceil < 0{
        return -1
    }
    return ceil
}

/* Testing Code */
func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
    t := CreateBinarySearchTree(arr)
    fmt.Println(t.CeilBST(8))
}