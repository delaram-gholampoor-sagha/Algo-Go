package main

import (
	"fmt"
)

func GeneratePermutations(Movies []string) [][]string {

	Size := len(Movies)
	var Backtrack func(int)
	var Output [][]string

	Backtrack = func(First int) {

		/* If all integers of given array `Movies` are used and
		   and Backtracking is performed add the permutations to Output array. */
		if First == Size {
			temp := make([]string, len(Movies[:]))
			copy(temp, Movies[:])
			Output = append(Output, temp)
		}

		/* Perform Backtracking for the Size of a given array. */
		for i := First; i < Size; i++ {

			/* Swap: In the current permutaion place i-th integer first. */
			Movies[First], Movies[i] = Movies[i], Movies[First]
			/* Complete permuations using the next integers. */
			Backtrack(First + 1)
			/* Swap and Backtrack */
			Movies[First], Movies[i] = Movies[i], Movies[First]
		}
	}
	Backtrack(0)
	return Output
}

func main() {

	// Example #1
	Input := []string{"Frozen", "Dune", "Coco"}
	Output := GeneratePermutations(Input)
	fmt.Println("Output 1:", Output)

	// Example #2
	Input = []string{"Frozen", "Dune", "Coco", "Melificient"}
	Output = GeneratePermutations(Input)
	fmt.Println("Output 2:", Output)

	// Example #3
	Input = []string{"Dune", "Coco"}
	Output = GeneratePermutations(Input)
	fmt.Println("Output 3:", Output)
}
