package main

import "fmt"

type node struct {
	nvalue *Node
	next   *node
}

type Stack struct {
	head *node
	size int
}

func (s *Stack) Length() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Peek() (*Node, bool) {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return nil, false
	}
	return s.head.nvalue, true
}

func (s *Stack) Push(nvalue *Node) {
	s.head = &node{&Node{nvalue.value, nvalue.left, nvalue.right}, s.head}
	s.size++
}

func (s *Stack) Pop() *Node {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return nil
	}
	temp := new(Node)
	temp.value = s.head.nvalue.value
	temp.left = s.head.nvalue.left
	temp.right = s.head.nvalue.right
	s.head = s.head.next
	s.size--
	return temp
}
