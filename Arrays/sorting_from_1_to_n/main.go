package main

import "fmt"

func Sort1toN(arr []int, size int) {
	curr, value, next := 0, 0, 0
	for i := 0; i < size; i++ {
		curr = i
		value = -1
		/* swaps to move elements in the proper position. */
		for curr >= 0 && curr < size && arr[curr] != curr+1 {
			next = arr[curr]
			arr[curr] = value
			value = next
			curr = next - 1

		}
	}
}

// Testing code
func main() {
	arr := []int{5, 3, 2, 1, 4}
	size := len(arr)
	Sort1toN(arr, size)
	fmt.Println(arr)
}
