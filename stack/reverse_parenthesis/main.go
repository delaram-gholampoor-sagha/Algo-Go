package main
import "fmt"
import "math"

func reverseParenthesis(expn string, size int) int {

  stk := new(Stack)
  openCount := 0
  closeCount := 0
  var ch rune

  if (size % 2 == 1) {
    fmt.Println("Invalid odd length " , size)
    return -1
  }

  for i := 0; i < size; i++ {
    ch = rune(expn[i])
    if (ch == '(') {
      stk.Push(ch)
    } else if (ch == ')') {
      if (stk.Length() != 0 && stk.Top() == '(') {
        stk.Pop()
      } else {
        stk.Push(ch)
      }
    }
  }

  for (stk.Length() != 0) {
    if (stk.Top() == '(') {
      openCount += 1
    } else {
      closeCount += 1
    }
    stk.Pop()
  }

  reversal := int(math.Ceil(float64(openCount) / 2.0)) +
                  int(math.Ceil(float64(closeCount) / 2.0))
  return reversal
}

// Testing code
func main() {
  expn := ")(())((("
  size := len(expn)
  value :=reverseParenthesis(expn, size)
  fmt.Println("Given expn : " , expn)
  fmt.Println("reverse Parenthesis is : " , value)
}