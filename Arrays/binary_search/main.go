package main

import "fmt"

func BinarySearch(data []int, value int) bool {
	var mid int
	low := 0
	high := len(data) - 1

	for low <= high {
		mid = (low + high) / 2
		if data[mid] == value { // values find
			return true
		} else {
			if data[mid] < value { // move to left
				low = mid + 1
			} else {
				high = mid - 1 // move to right
			}
		}
	}
	return false // value not found
}

// Testing code
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println("BinarySearch:", BinarySearch(arr, 8))
	fmt.Println("BinarySearch:", BinarySearch(arr, 3))
}
