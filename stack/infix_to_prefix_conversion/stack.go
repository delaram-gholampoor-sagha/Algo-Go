package main

import "fmt"

type Stack struct {
	s []int
}

func (s *Stack) IsEmpty() bool {
	length := len(s.s)
	return length == 0
}

func (s *Stack) Length() int {
	length := len(s.s)
	return length
}

func (s *Stack) Print() {
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}

func (s *Stack) Push(value int) {
	s.s = append(s.s, value)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() == true {
		fmt.Print("Stack in empty.")
		return 0
	}
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

func (s *Stack) Top() int {
	if s.IsEmpty() == true {
		fmt.Print("Stack in empty.")
		return 0
	}
	length := len(s.s)
	res := s.s[length-1]
	return res
}

func precedence(x rune) int {
	if x == '(' {
		return (0)
	}
	if x == '+' || x == '-' {
		return (1)
	}
	if x == '*' || x == '/' || x == '%' {
		return (2)
	}
	if x == '^' {
		return (3)
	}
	return (4)
}

func InfixToPostfix(expn string) string {
	stk := new(Stack)
	output := ""

	for _, ch := range expn {
		if ch <= '9' && ch >= '0' {
			output = output + string(ch)
		} else {
			switch ch {
			case '+', '-', '*', '/', '%', '^':
				for stk.IsEmpty() == false && precedence(ch) <= precedence(rune(stk.Top())) {
					out := rune(stk.Pop())
					output = output + " " + string(out)
				}
				stk.Push(int(ch))
				output = output + " "
			case '(':
				stk.Push(int(ch))
			case ')':
				for stk.IsEmpty() == false && stk.Top() != '(' {
					out := rune(stk.Pop())
					output = output + " " + string(out) + " "
				}
				if stk.IsEmpty() == false && stk.Top() == '(' {
					stk.Pop()
				}
			}
		}
	}

	for stk.IsEmpty() == false {
		out := rune(stk.Pop())
		output = output + string(out)
	}
	return output
}
