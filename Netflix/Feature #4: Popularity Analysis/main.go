package main

import (
	"fmt"
)

func identifyTitles(scores []int) bool {

	increasing := true
	decreasing := true

	for i := 0; i < len(scores)-1; i++ {
		if scores[i] > scores[i+1] {
			increasing = false
		}
		if scores[i] < scores[i+1] {
			decreasing = false
		}
	}

	return increasing || decreasing
}

func main() {
	movieRatings := [][]int{{1, 2, 2, 3}, {4, 5, 6, 3, 4}, {8, 8, 7, 6, 5, 4, 4, 1}}

	for _, movieRating := range movieRatings {
		if identifyTitles(movieRating) {
			fmt.Println("Title Identified and Separated")
		} else {
			fmt.Println("Title Score Fluctuating")
		}
	}
}
