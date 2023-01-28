package main
import "fmt"

func findDuplicateParenthesis(expn string, size int) bool {
	stk := new(Stack)
	var ch byte
	var count int

	for i := 0; i < size; i++ {
		ch = expn[i]
		if (ch == ')') {
			count = 0
			for (stk.Length() != 0 && byte(stk.Top()) != '(') {
				stk.Pop()
				count += 1
			}
			if (count <= 1) {
				return true
			}
		} else {
			stk.Push(int(ch))
		}
	}
	return false
}

// Testing code
func main() {
	// expn = "(((a+(b))+(c+d)))"
	// expn = "(b)"
	expn := "(((a+b))+c)"
	fmt.Println("Given expn : " , expn)
	size := len(expn)
	value := findDuplicateParenthesis(expn, size)
	fmt.Println("Duplicate Parenthesis Found : " , value)
}