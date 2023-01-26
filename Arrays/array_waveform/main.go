package main

import "fmt"

func swap(data []int, x int, y int) {
	data[x], data[y] = data[y], data[x]
}

func WaveArray(arr []int) {
	size := len(arr)
	/* Odd elements are lesser than even elements. */
	for i := 1; i < size; i += 2 {
		if (i-1) >= 0 && arr[i] > arr[i-1] { //swap with left neighbour
			swap(arr, i, i-1)
		}
		if (i+1) < size && arr[i] > arr[i+1] { //swap with right neighbour
			swap(arr, i, i+1)
		}
	}
}

/* Testing code */
func main() {
	arr := []int{8, 1, 2, 3, 4, 5}
	WaveArray(arr)
	fmt.Println(arr)
}
