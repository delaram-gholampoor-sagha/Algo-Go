package main

import (
	"fmt"
)

/*
	Declare a FreqStack struct containing Frequency and Group hashmaps

and MaxFrequency inetegr
*/
type FreqStack struct {
	Frequency    map[string]int
	Group        map[int][]string
	MaxFrequency int
}

// Use constructor to initialize the FreqStack object
func Constructor() FreqStack {
	return FreqStack{
		Frequency:    make(map[string]int),
		Group:        make(map[int][]string),
		MaxFrequency: 0,
	}
}

// Use Push function to push the ShowName into the FreqStack
func (this *FreqStack) Push(ShowName string) {
	// Get the frequency for the given ShowName
	f, _ := this.Frequency[ShowName]
	// Increment the frequency for the given ShowName
	f += 1
	this.Frequency[ShowName] = f
	/* Check if the maximum frequency is lower that the new frequency
	   of the given show */
	if f > this.MaxFrequency {
		this.MaxFrequency = f
	}
	// Save the given ShowName for the new calculated frequency
	this.Group[f] = append(this.Group[f], ShowName)
}

// Use the Pop function to pop the ShowName from the FreqStack
func (this *FreqStack) Pop() string {
	// Initialize an empty show string
	Show := ""
	// Check if there is any show present in a stack
	if this.MaxFrequency > 0 {
		// Fetch the top of the Group[this.MaxFrequency] stack
		Show = this.Group[this.MaxFrequency][len(this.Group[this.MaxFrequency])-1]
		// Pop the top of the Group[this.MaxFrequency] stack
		this.Group[this.MaxFrequency] = this.Group[this.MaxFrequency][:len(this.Group[this.MaxFrequency])-1]
		// Decrement the frequency after the show has been popped
		this.Frequency[Show] -= 1
		if len(this.Group[this.MaxFrequency]) == 0 {
			this.MaxFrequency -= 1
		}
	}
	return Show
}

// Driver code
func main() {
	obj := Constructor()
	obj.Push("Queen's Gambit")
	obj.Push("Teen Wolf")
	obj.Push("Queen's Gambit")
	fmt.Println("...User navigates to the browser...")
	fmt.Println("Continue Watching :", obj.Pop())
	fmt.Println()
	obj.Push("Teen Wolf")
	obj.Push("Bigderton")
	fmt.Println("...User navigates to the browser...")
	fmt.Println("Continue Watching :", obj.Pop())
	fmt.Println()
	obj.Push("Queen's Gambit")
	obj.Push("Teen Wolf")
	obj.Push("Bigderton")
	for i := 0; i < 7; i++ {
		fmt.Println("...User navigates to the browser...")
		fmt.Println("Continue Watching :", obj.Pop())
		fmt.Println()
	}

}
