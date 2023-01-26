package main

import (
	"fmt"
)

func SequentialSearch(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		}
	}
	return false
}

// Testing code
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println("SequentialSearch:", SequentialSearch(arr, 8))
	fmt.Println("SequentialSearch:", SequentialSearch(arr, 9))
}
