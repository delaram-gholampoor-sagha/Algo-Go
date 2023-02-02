package main

import (
	"container/heap"
	"fmt"
)

type MedianOfAges struct {
	maxHeap *MaxHeap //containing first half of numbers
	minHeap *MinHeap //containing second half of numbers
}

// Initialize the heaps
func new() *MedianOfAges {
	min := &MinHeap{}
	max := &MaxHeap{}
	heap.Init(min)
	heap.Init(max)
	return &MedianOfAges{minHeap: min, maxHeap: max}
}

func (med *MedianOfAges) FindMedian() float64 {
	if med.maxHeap.Len() == med.minHeap.Len() {
		// we have even number of elements, take the average of middle two elements
		return float64(float64(med.maxHeap.Top()+med.minHeap.Top()) / 2.0)
	}
	// because max-heap will have one more element than the min-heap
	return float64(med.maxHeap.Top())
}

func (med *MedianOfAges) InsertNum(num int) {
	if med.maxHeap.Empty() || med.maxHeap.Top() >= num {
		heap.Push(med.maxHeap, num)
	} else {
		heap.Push(med.minHeap, num)
	}

	// either both the heaps will have equal number of elements or max-heap will have one
	// more element than the min-heap
	if med.maxHeap.Len() > med.minHeap.Len()+1 {
		heap.Push(med.minHeap, heap.Pop(med.maxHeap).(int))
	} else if med.maxHeap.Len() < med.minHeap.Len() {
		heap.Push(med.maxHeap, heap.Pop(med.minHeap).(int))
	}
}

func main() {

	medianOfAges := new()
	medianOfAges.InsertNum(22)
	medianOfAges.InsertNum(35)
	fmt.Printf("The recommended content will be for ages under: %f\n", medianOfAges.FindMedian())
	medianOfAges.InsertNum(30)
	fmt.Printf("The recommended content will be for ages under: %f\n", medianOfAges.FindMedian())
	medianOfAges.InsertNum(25)
	fmt.Printf("The recommended content will be for ages under: %f\n", medianOfAges.FindMedian())
}
