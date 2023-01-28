package main
import "fmt"

func printParenthesisNumber(expn string, size int) {
	var ch byte
	stk := new(Stack)
	output := ""
	var count int = 1
	for i := 0; i < size; i++ {
		ch = expn[i]
		if (ch == '(') {
			stk.Push(count)
			output += fmt.Sprintf("%v",count)
			count += 1
		} else if (ch == ')') {
			output += fmt.Sprintf("%v",stk.Pop())
		}
	}
	fmt.Println("Parenthesis Count : " , output)
}

//Testing code
func main() {
	expn1 := "(((a+(b))+(c+d)))"
	size := len(expn1)
	fmt.Println("Given expn : " , expn1)
	printParenthesisNumber(expn1, size)

	expn2 := "(((a+b))+c)((("
	size = len(expn2)
	fmt.Println("Given expn : " , expn2)
	printParenthesisNumber(expn2, size)
}