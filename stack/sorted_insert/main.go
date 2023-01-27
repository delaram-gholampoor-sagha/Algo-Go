package main

func sortedInsert(stk *Stack, element int) {
	var temp int
	if stk.Length() == 0 || element > stk.Top() {
		stk.Push(element)
	} else {
		temp = stk.Pop()
		sortedInsert(stk, element)
		stk.Push(temp)
	}
}

// Testing code
func main() {
	stk := new(Stack)
	stk.Push(1)
	stk.Push(2)
	stk.Push(4)
	stk.Push(9)
	sortedInsert(stk, 3)
	sortedInsert(stk, 7)
	sortedInsert(stk, 8)
	stk.Print()
}
