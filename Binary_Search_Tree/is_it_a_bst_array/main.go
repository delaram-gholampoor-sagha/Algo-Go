package main
import "fmt"

func isBSTArray( preorder[] int, size int) bool {
    stk := new(Stack)
    var value int
    root := -999999;
    for i := 0; i < size; i++ {
        value = preorder[i]
        // If the value of the right child is less than the root.
        if (value < root){
            return false
        }
        // First left child values will be popped
        // Last popped value will be the root.

        for (stk.Length() > 0 && stk.Top() < value){
            root = stk.Pop()
        }
        // add current value to the stack.
        stk.Push(value)
    }
    return true
}

/* Testing Code */
func main() {
    arr3 := []int{5, 2, 4, 6, 9, 10}
    fmt.Println(isBSTArray(arr3, len(arr3)))
    
    arr2 := []int{5, 2, 6, 4, 7, 9, 10}
    fmt.Println(isBSTArray(arr2, len(arr2)))
}