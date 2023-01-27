package main

func sortStack2(stk *Stack) {
    var temp int
    stk2 := new(Stack)
    for (stk.Length() > 0) {
        temp = stk.Pop()
        for ((stk2.Length() > 0) && (stk2.Top() < temp)) {
            stk.Push(stk2.Pop())
        }
        stk2.Push(temp)
    }
    for (stk2.Length() > 0) {
        stk.Push(stk2.Pop())
    }
}

// Testing code
func main() {
    stk := new(Stack)
    stk.Push(1)
    stk.Push(4)
    stk.Push(3)
    stk.Push(2)
    sortStack2(stk)
    stk.Print()
}