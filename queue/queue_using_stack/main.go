package main
import "fmt"

type QueueUsingStack struct {
  stk1 Stack
  stk2 Stack
}

func (que *QueueUsingStack) Add(value int) {
    que.stk1.Push(value)
}

func (que *QueueUsingStack) Remove() int {
  var value int
  if que.stk2.IsEmpty() == false {
    value = que.stk2.Pop()
    return value
  }
  for que.stk1.IsEmpty() == false {
    value = que.stk1.Pop()
    que.stk2.Push(value)
  }
  value = que.stk2.Pop()
  return value
}

func (que *QueueUsingStack) Length() int {
  return (que.stk1.Length() + que.stk2.Length());
}

func (que *QueueUsingStack) IsEmpty() bool {
  return (que.stk1.Length() + que.stk2.Length()) == 0;
}

// Testing Code
func main() {
  que := new(QueueUsingStack)
  que.Add(1)
  que.Add(2)
  que.Add(3)
  fmt.Println(que.Remove())
  fmt.Println(que.Remove())
  fmt.Println(que.Remove())
}