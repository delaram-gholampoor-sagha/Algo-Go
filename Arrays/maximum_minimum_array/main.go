package main

import (
	"fmt"
)

func MaxMinArr(arr []int, size int) {
	aux := make([]int, size)
	copy(aux, arr)
	start := 0
	stop := size - 1
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			arr[i] = aux[stop]
			stop -= 1
		} else {
			arr[i] = aux[start]
			start += 1
		}
	}
}

/* Testing code */
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	size := len(arr)
	MaxMinArr(arr, size)
	fmt.Println(arr)
}
