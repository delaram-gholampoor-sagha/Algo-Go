package main

import "fmt"

func IsBalancedParenthesis(expn string) bool {
	stk := new(Stack)
	for _, ch := range expn {
		switch ch {
		case '{', '[', '(':
			stk.Push(ch)
		case '}':
			val := stk.Pop()
			if val != '{' {
				return false
			}
		case ']':
			val := stk.Pop()
			if val != '[' {
				return false
			}
		case ')':
			val := stk.Pop()
			if val != '(' {
				return false
			}
		}
	}
	return stk.IsEmpty()
}

// Testing code
func main() {
	expn := "{()}]"
	value := IsBalancedParenthesis(expn)
	fmt.Println("Given Expn :", expn)
	fmt.Println("IsBalancedParenthesis :", value)
}
