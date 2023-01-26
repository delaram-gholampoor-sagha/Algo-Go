package main

import "fmt"

func MaxSubArraySum(data []int) int {
	size := len(data)
	maxSoFar := 0
	maxEndingHere := 0
	for i := 0; i < size; i++ {
		maxEndingHere = maxEndingHere + data[i] //adding the values till current value
		if maxEndingHere < 0 {                  // reset index if its negative
			maxEndingHere = 0
		}

		//reset maximum global if its smaller then current one
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
	}
	return maxSoFar
}

// Testing code
func main() {
	data := []int{1, -2, 3, 4, -4, 6, -14, 6, 2}
	fmt.Println("Max sub array sum :", MaxSubArraySum(data))
}
