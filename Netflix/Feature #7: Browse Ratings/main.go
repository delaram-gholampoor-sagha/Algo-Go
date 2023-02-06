package main

import (
	"fmt"
)

type MaxStack struct {
	maxSize int
	//We will use two stacks mainStack to hold original values
	//and maxStack to hold maximum values. Top of maxStack will always
	//be the maximum value from mainStack
	mainStack    Stack
	maximumStack Stack
}

// removes and returns value from stack
func (m *MaxStack) Pop() int {
	//1. Pop element from maxStack to make it sync with mainStack,
	//2. Pop element from mainStack and return that value
	m.maximumStack, _ = m.maximumStack.Pop()
	top := m.mainStack.Top()
	m.mainStack, _ = m.mainStack.Pop()
	return top
}

// Pushes value into the stack
func (m *MaxStack) Push(value int) {
	//1. Push value in mainStack and check value with the top value of maxStack
	//2. If value is greater than top, then Push top in maxStack
	//else Push value in maxStack
	m.mainStack = m.mainStack.Push(value)
	if !m.maximumStack.Empty() && m.maximumStack.Top() > value {
		m.maximumStack = m.maximumStack.Push(m.maximumStack.Top())
	} else {
		m.maximumStack = m.maximumStack.Push(value)
	}
}

// returns maximum value in O(1)
func (m *MaxStack) MaxRating() int {
	return m.maximumStack.Top()
}

func main() {

	stack := &MaxStack{maxSize: 7}
	stack.Push(5)
	stack.Push(0)
	stack.Push(2)
	stack.Push(4)
	stack.Push(6)
	stack.Push(3)
	stack.Push(10)

	fmt.Println("Maximum Rating:", stack.MaxRating())

	stack.Pop()
	fmt.Println("\nAfter clicking back button")

	fmt.Println("\nMaximum Rating: ", stack.MaxRating())
}
