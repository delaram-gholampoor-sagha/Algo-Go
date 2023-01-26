package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Testing code
func main() {
	fmt.Println("Desired number in fibonacci series :", fibonacci(10))
}
