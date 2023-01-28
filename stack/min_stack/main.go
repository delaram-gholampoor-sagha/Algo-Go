package main
import "fmt"

type MinStack struct {
    maxSize int
	//We will use two stacks mainStack to hold original values
	//and minStack to hold minimum values. Top of minStack will always
	//be the minimum value from mainStack
    mainStack Stack
	minimumStack Stack
}

//removes and returns value from stack
func (m *MinStack) Pop() int{
	//1. Pop element from minStack to make it sync with mainStack,
	//2. Pop element from mainStack and return that value
	_ = m.minimumStack.Pop()
	top := m.mainStack.Top()
	_ = m.mainStack.Pop()
	return top
}

//Pushes value into the stack
func (m *MinStack) Push(value int){
	//1. Push value in mainStack and check value with the top value of minStack
	//2. If value is smaller than top, then Push top in minStack
	//else Push value in minStack
	m.mainStack.Push(value)
	if !m.minimumStack.IsEmpty() && m.minimumStack.Top() < value {
		m.minimumStack.Push(m.minimumStack.Top())
	} else {
		m.minimumStack.Push(value)
	}
}

//returns minimum value in O(1)
func (m *MinStack) Min() int{
	return m.minimumStack.Top()
}

func main() {

    stack := &MinStack{maxSize: 7}
    stack.Push(5)
    stack.Push(9)
    stack.Push(2)
    stack.Push(4)
    stack.Push(6)
    stack.Push(3)
    stack.Push(1)

    fmt.Println("Minimum:", stack.Min())

    stack.Pop()
    fmt.Println("\nAfter clicking back button")

    fmt.Println("\nMinimum: ", stack.Min())
}