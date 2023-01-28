package main
import "fmt"

func reverseKElementInQueue(que *Queue, k int) {
    stk := new(Stack)
    i := 0
    var diff, temp int
    for (que.Length() != 0 && i < k) {
        stk.Push(que.Dequeue().(int))
        i++
    }
    for (stk.Length() > 0) {
        que.Enqueue(stk.Pop())
    }
    diff = que.Length() - k
    for (diff > 0) {
        temp = que.Dequeue().(int)
        que.Enqueue(temp)
        diff -= 1
    }
}

// Testing code
func main() {
    que := new(Queue)
    que.Enqueue(1)
    que.Enqueue(2)
    que.Enqueue(3)
    que.Enqueue(4)
    fmt.Print("Queue before Reversal: ")
    que.Print()
    fmt.Print("Queue after Reversal: ")
    reverseKElementInQueue(que, 2)
    que.Print()
}