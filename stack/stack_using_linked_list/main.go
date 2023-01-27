package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type StackLinkedList struct {
	head *Node
	size int
}

// Size() function will return the size of the linked list.
func (s *StackLinkedList) Size() int {
	return s.size
}

/*
	IsEmpty() function will return true if size of the linked list is

equal to zero or false in all other cases.
*/
func (s *StackLinkedList) IsEmpty() bool {
	return s.size == 0
}

/*First, the Peek() function will check if the stack is empty.
If not, it will return the peek value of the stack, i.e., it will return
the head value of the linked list. */

func (s *StackLinkedList) Peek() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return 0, false
	}
	return s.head.value, true
}

// Push() function  will append the value to the linked list.
func (s *StackLinkedList) Push(value int) {
	s.head = &Node{value, s.head}
	s.size++
}

/*
In the pop() function, first it will check that the stack is not empty.
Then it will pop the value from the linked list and return it.
*/
func (s *StackLinkedList) Pop() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return 0, false
	}
	value := s.head.value
	s.head = s.head.next
	s.size--
	return value, true
}

/* Print() function will print the elements of the linked list. */
func (s *StackLinkedList) Print() {
	temp := s.head
	fmt.Print("Value stored in stock are :: ")
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

// Testing code
func main() {
	s := new(StackLinkedList)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	val, _ := s.Peek()
	fmt.Println("Peek() of a stack is:", val)
	fmt.Print("Stack consist following elements: ")
	for s.IsEmpty() == false {
		val, _ = s.Pop()
		fmt.Print(val, " ")
	}
}
