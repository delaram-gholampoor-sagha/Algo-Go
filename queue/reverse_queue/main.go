package main
import "fmt"

func reverseQueue(que *Queue) {
  stk := new(Stack)
  for (que.Length() != 0) {
    stk.Push(que.Dequeue().(int))
  }
  for (stk.Length() > 0) {
    que.Enqueue(stk.Pop())
  }
}

// Testing code
func main() {
  que := new(Queue)
  que.Enqueue(1)
  que.Enqueue(2)
  que.Enqueue(3)
  que.Enqueue(4)
  fmt.Print("Queue before reversal : ")
  que.Print()
  reverseQueue(que)
  fmt.Print("Queue after reversal : ")
  que.Print()
}