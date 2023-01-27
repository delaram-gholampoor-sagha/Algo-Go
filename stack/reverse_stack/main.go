package main
import "fmt"

func reverseStack2(stk *Stack) {
    que := new(Queue)
    for (stk.Length() > 0) {
        que.Enqueue(stk.Pop())
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
    stk.Push(3)
    fmt.Print("Stack before reversal : ")
    stk.Print()
    reverseStack2(stk)
    fmt.Print("Stack after reversal : ")
    stk.Print()
}