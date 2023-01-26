package main

import (
	"fmt"
)

func towerOfHanoi(num int, src byte, dst byte, temp byte) {
	if num < 1 {
		return
	}
	towerOfHanoi(num-1, src, temp, dst)
	fmt.Printf("Move disk %d from peg %c to peg %c \n", num, src, dst)
	towerOfHanoi(num-1, temp, dst, src)
}

// Testing code
func main() {
	num := 3
	fmt.Println("The sequence of moves involved in the Tower of Hanoi are :")
	towerOfHanoi(num, 'A', 'C', 'B')
}
