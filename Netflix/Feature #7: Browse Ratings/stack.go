package main

type Stack []int

func (s Stack) Push(v int) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, int) {

	l := len(s)
	if l == 0 {
		return s, 0
	}
	return s[:l-1], s[l-1]
}

func (s Stack) Top() int {
	l := len(s)
	return s[l-1]
}

func (s Stack) Empty() bool {
	return len(s) == 0
}
