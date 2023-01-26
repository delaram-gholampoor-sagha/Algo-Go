package main

import "fmt"

func RotateArray(data []int, k int) {
	n := len(data)
	// first part
	ReverseArray(data, 0, k-1) // reverse 1st part of array till kth position
	ReverseArray(data, k, n-1) // reverse from kth position till end

	//second part
	ReverseArray(data, 0, n-1) // reverse whole array
}

func ReverseArray(data []int, start int, end int) {
	i := start
	j := end
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}

// Testing code
func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Input array :", arr)
	RotateArray(arr, 2)
	fmt.Println("Rotated array :", arr)
}
