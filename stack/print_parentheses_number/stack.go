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
    if s.IsEmpty() == true{
        fmt.Print("Stack in empty.")
        return 0
    }
    length := len(s.s)
    res := s.s[length-1]
    s.s = s.s[:length-1]
    return res
}

func (s *Stack) Top() int {
    if s.IsEmpty() == true{
        fmt.Print("Stack in empty.")
        return 0
    }
    length := len(s.s)
    res := s.s[length-1]
    return res
}