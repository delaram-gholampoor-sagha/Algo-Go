package main

import (
	"fmt"
	"math/rand"
)

type LinkedListNode struct {
	key              int
	data             int
	next             *LinkedListNode
	arbitrartPointer *LinkedListNode
}

func createLinkedList(lst []int) *LinkedListNode {
	var head *LinkedListNode
	var tail *LinkedListNode
	for _, x := range lst {
		newNode := &LinkedListNode{data: x}
		if head == nil {
			head = newNode
		} else {
			tail.next = newNode
		}
		tail = newNode
	}
	return head
}

func insertAtHead(head *LinkedListNode, data int) *LinkedListNode {
	newNode := &LinkedListNode{data: data}
	newNode.next = head
	return newNode
}

func insertAtTail(head *LinkedListNode, data int) *LinkedListNode {
	newNode := &LinkedListNode{data: data}
	if head == nil {
		return newNode
	}

	temp := head

	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
	return head
}

func createRandomList(length int) *LinkedListNode {
	var listHead *LinkedListNode
	for i := 0; i < length; i++ {
		listHead = insertAtHead(listHead, rand.Intn(100))
	}
	return listHead
}

func toList(head *LinkedListNode) []int {
	var lst []int
	temp := head
	for temp != nil {
		lst = append(lst, temp.data)
		temp = temp.next
	}
	return lst
}

func display(head *LinkedListNode) {
	temp := head
	for temp != nil {
		fmt.Printf("%d", temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("\n")
}

func mergeAlternating(list1, list2 *LinkedListNode) *LinkedListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	head := list1

	for list1.next != nil && list2 != nil {
		temp := list2
		list2 = list2.next

		temp.next = list1.next
		list1.next = temp
		list1 = temp.next
	}

	if list1.next == nil {
		list1.next = list2
	}

	return head
}

func isEqual(list1, list2 *LinkedListNode) bool {
	if list1 == list2 {
		return true
	}

	for list1 != nil && list2 != nil {
		if list1.data != list2.data {
			return false
		}

		list1 = list1.next
		list2 = list2.next
	}

	return (list1 == list2)
}
