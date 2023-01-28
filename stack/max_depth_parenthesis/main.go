package main
import "fmt"

func maxDepthParenthesis(expn string, size int) int {
    stk := new(Stack)
    maxDepth := 0
    depth := 0
    var ch byte
    for i := 0; i < size; i++ {
        ch = expn[i]
        if (ch == '(') {
            stk.Push(ch)
            depth += 1
        } else if (ch == ')') {
            stk.Pop()
            depth -= 1
        }
        if (depth > maxDepth) {
            maxDepth = depth
        }
    }
    return maxDepth
}

// Testing code
func main() {
    expn := "((((A)))((((BBB()))))()()()())"
    size := len(expn)
    value :=maxDepthParenthesis(expn, size)
    fmt.Println("Given expn " , expn)
    fmt.Println("Max depth parenthesis is " , value)
}