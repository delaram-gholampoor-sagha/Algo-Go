package main

func bottomInsert(stk *Stack, element int) {
	var temp int
	if stk.Length() == 0 {
		stk.Push(element)
	} else {
		temp = stk.Pop()
		bottomInsert(stk, element)
		stk.Push(temp)
	}
}

// Testing code
func main() {
	stk := new(Stack)
	stk.Push(1)
	stk.Push(2)
	stk.Push(3)
	bottomInsert(stk, 5)
	stk.Print()
}
