package main
import "fmt"

const capacity = 100

type Queue struct {
  size int
  data [capacity]interface{}
  front int
  back int
}

func (q *Queue) Enqueue(value interface{}) {
  if q.size >= capacity {
    return
  }
  q.size++
  q.data[q.back] = value
  q.back = (q.back + 1) % (capacity - 1)
}

func (q *Queue) Dequeue() interface{} {
  var value interface{}
  if q.size <= 0 {
    return 0
  }
  q.size--
  value = q.data[q.front]
  q.front = (q.front + 1) % (capacity - 1)
  return value
}

func (q *Queue) IsEmpty() bool {
  return q.size == 0
}

func (q *Queue) Length() int {
  return q.size
}

//Testing Code
func main() {
  q := new(Queue)
  q.Enqueue(1)
  q.Enqueue(2)
  q.Enqueue(3)
  for q.IsEmpty() == false {
  val, _ := q.Dequeue().(int)
  fmt.Println(val, " ")
  }
}