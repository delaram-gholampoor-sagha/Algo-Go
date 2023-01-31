package main


func CreateBinarySearchTree(arr []int) *Tree {
    t := new(Tree)
    size := len(arr)
    t.root = createBinarySearchTreeUtil(arr, 0, size-1)
    return t
}

func createBinarySearchTreeUtil(arr []int, start int, end int) *Node {
    if start > end {
        return nil
    }
    mid := (start + end) / 2
    curr := new(Node)
    curr.value = arr[mid]
    curr.left = createBinarySearchTreeUtil(arr, start, mid-1)
    curr.right = createBinarySearchTreeUtil(arr, mid+1, end)
    return curr
}

/* Testing Code */
func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    t := CreateBinarySearchTree(arr)
    t.PrintInOrder()
}