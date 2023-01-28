package main
import "fmt"

type Node struct {
    value int
    next *Node
    prev *Node
}

type Deque struct {
    head *Node
    tail *Node
    size int
}

func (q *Deque) Size() int {
    return q.size
}

func (q *Deque) IsEmpty() bool {
    return q.size == 0
}

func (q *Deque) PeekFront() int {
    if q.IsEmpty() {
        fmt.Println("QueueEmptyException")
        return 0
    }
    return q.head.value
}

func (q *Deque) PeekBack() int {
    if q.IsEmpty() {
        fmt.Println("QueueEmptyException")
        return 0
    }
    return q.tail.value
}

func (q *Deque) AddBack(value int) {
    temp := &Node{value, nil, nil}
    if q.head == nil {
        q.head = temp
        q.tail = temp
    } else {
        q.tail.next = temp
        temp.prev = q.tail
        q.tail = temp
    }
    q.size++
}

func (q *Deque) AddFront(value int) {
    temp := &Node{value, nil, nil}
    if q.head == nil {
        q.head = temp
        q.tail = temp
    } else {
        temp.next = q.head
        q.head.prev = temp
        q.head = temp
    }
    q.size++
}

func (q *Deque) RemoveFront() int {
    if q.IsEmpty() {
        fmt.Println("QueueEmptyException")
        return 0
    }
    q.size--
    value := q.head.value
    q.head = q.head.next

    if q.head != nil {
        q.head.prev = nil
    } else {
        q.tail = nil
    }

    return value
}

func (q *Deque) RemoveBack() int {
    if q.IsEmpty() {
        fmt.Println("QueueEmptyException")
        return 0
    }
    q.size--
    value := q.tail.value
    q.tail = q.tail.prev
    if q.tail != nil {
        q.tail.next = nil
    } else {
        q.head = nil
    }
    return value
}

func (q *Deque) Print() {
    temp := q.head
    for temp != nil {
        fmt.Print(temp.value, " ")
        temp = temp.next
    }
    fmt.Println()
}

// Testing Code
func main() {
    q := new(Deque)
    q.AddFront(1)
    q.AddFront(1)
    q.AddBack(2)
    q.AddBack(2)
    q.Print()
    fmt.Println(q.RemoveBack())
    q.Print()
    fmt.Println(q.RemoveFront())
    q.Print()
}