package main

import "fmt"

type StackInt struct {
	s []int
}

/* Other Methods */

func (s *StackInt) IsEmpty() bool {
	length := len(s.s)
	return length == 0
}

func (s *StackInt) Length() int {
	length := len(s.s)
	return length
}

func (s *StackInt) Print() {
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}

func (s *StackInt) Push(value int) {
	s.s = append(s.s, value)
}

func (s *StackInt) Pop() int {
	if s.IsEmpty() == true {
		fmt.Print("Stack is empty.")
		return 0
	}
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

func (s *StackInt) Top() int {
	if s.IsEmpty() == true {
		fmt.Print("Stack is empty.")
		return 0
	}
	length := len(s.s)
	res := s.s[length-1]
	return res
}

// Test code for stack.
func main() {
	stack := new(StackInt)
	//Push element to the stack
	stack.Push(6)
	stack.Push(3)
	stack.Push(2)
	stack.Push(5)
	//Retrieve top element from the stack
	fmt.Println("Top() of the stack is: ", stack.Top())
	//Pop elements from the stack
	fmt.Print("Stack consist of following elements: ")
	for stack.IsEmpty() == false {
		fmt.Print(stack.Pop(), " ")
	}
}
