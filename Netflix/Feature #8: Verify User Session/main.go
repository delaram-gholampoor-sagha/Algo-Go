package main

import (
	"fmt"
)

func verifySession(pushOp, popOp []int) bool {

	M := len(pushOp)
	N := len(popOp)
	if M != N {
		return false
	}

	// var stack Stack;
	var stack Stack;

	i := 0
	for _, pid := range pushOp {
		stack = stack.Push(pid)
		for !stack.Empty() && stack.Top() == popOp[i] {
			stack, _ = stack.Pop()
			i++
		}
	}

	return stack.Empty()
}

func main() {
	// Driver code
	pushOp := []int{1, 2, 3, 4, 5}
	popOp := []int{5, 4, 3, 2, 1}

	if verifySession(pushOp, popOp) {
		fmt.Println("Session Successfull!")
	} else {
		fmt.Println("Session Faulty!")
	}

	pushOp = []int{6, 7, 8, 9, 10}
	popOp = []int{8, 10, 7, 9}

	if verifySession(pushOp, popOp) {
		fmt.Println("Session Successfull!")
	} else {
		fmt.Println("Session Faulty!")
	}
}
