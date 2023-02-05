package main

type LinkedListNode struct{
	key int
	val int
	freq int
	next *LinkedListNode
	prev *LinkedListNode
}


type LinkedList struct{ 
	head *LinkedListNode
	tail *LinkedListNode
}

func (l *LinkedList) Append(node *LinkedListNode){
	if l.head == nil{
		l.head = node
	} else{
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}

func (l *LinkedList) DeleteNode(node *LinkedListNode){
	if node.prev != nil {
		node.prev.next = node.next
	} else{
		l.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else{
		l.tail = node.prev
	}
	node.next = nil
	node.prev = nil
}