package main
import "fmt"

func reverseKElementInStack(stk *Stack, k int) {
    que := new(Queue)
    i := 0
    for (stk.Length() > 0 && i < k) {
        que.Enqueue(stk.Pop())
    i++
    }
    for (que.Length() != 0) {
        stk.Push(que.Dequeue().(int))
    }
}
// Testing code
func main() {
    stk := new(Stack)
    stk.Push(1)
    stk.Push(2)
    stk.Push(4)
    stk.Push(3)
    fmt.Print("Stack before reversal : ")
    stk.Print()
    reverseKElementInStack(stk, 2)
    fmt.Print("Stack after reversal : ")
    stk.Print()
}