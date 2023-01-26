package main

import (
	"fmt"
)

func swap(x, y int) (int, int) {
	return y, x
}

func Permutation(data []int, i int, length int) {
	if length == i {
		fmt.Printf("%v\n", data)
		return
	}
	for j := i; j < length; j++ {
		data[i], data[j] = swap(data[i], data[j])
		Permutation(data, i+1, length)
		data[i], data[j] = swap(data[i], data[j])
	}
}

// Testing code
func main() {
	var data [3]int
	for i := 0; i < 3; i++ {
		data[i] = i
	}
	Permutation(data[:], 0, 3)
}
