package main

import (
	"fmt"
)

func swap(x, y int) (int, int) {
	return y, x
}

func Partition01(arr []int, size int) {
	left := 0
	right := size - 1
	count := 0
	for left < right {
		for arr[left] == 0 {
			left += 1
		}
		for arr[right] == 1 {
			right -= 1
		}
		if left < right {
			arr[left], arr[right] = swap(arr[left], arr[right])
			count += 1
		}
	}

}

// Testing code
func main() {
	arr := []int{0, 1, 1, 0, 1, 0, 1, 1, 0, 0, 0, 1}
	Partition01(arr, len(arr))
	fmt.Println(arr)

}
