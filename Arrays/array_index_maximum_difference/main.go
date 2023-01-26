package main

import (
	"fmt"
)

func ArrayIndexMaxDiff2(arr []int, size int) int {
	leftMin := make([]int, size)
	rightMax := make([]int, size)
	leftMin[0] = arr[0]
	i, j := 0, 0
	var maxDiff = 0
	for i = 1; i < size; i++ {
		if leftMin[i-1] < arr[i] {
			leftMin[i] = leftMin[i-1]
		} else {
			leftMin[i] = arr[i]
		}
	}
	rightMax[size-1] = arr[size-1]
	for i = size - 2; i >= 0; i-- {
		if rightMax[i+1] > arr[i] {
			rightMax[i] = rightMax[i+1]
		} else {
			rightMax[i] = arr[i]
		}
	}
	i = 0
	j = 0
	maxDiff = -1
	for j < size && i < size {
		if leftMin[i] < rightMax[j] {
			if maxDiff < j-i {
				maxDiff = j - i
			}
			j = j + 1
		} else {
			i = i + 1
		}
	}
	return maxDiff
}

// Testing code
func main() {
	arr := []int{33, 9, 10, 3, 2, 60, 30, 33, 1}
	fmt.Println("ArrayIndexMaxDiff : ", ArrayIndexMaxDiff2(arr, len(arr)))
}
