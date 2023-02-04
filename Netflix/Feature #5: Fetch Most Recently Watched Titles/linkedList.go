package main

type LinkedListNode struct {
	key  int
	data int
	next *LinkedListNode
	prev *LinkedListNode
}

type LinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
	size int
}

func (l *LinkedList) InsertAtHead(key, data int) {

	newNode := &LinkedListNode{key: key, data: data}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}
	l.size++
}

func (l *LinkedList) InsertAtTail(key, data int) {
	newNode := &LinkedListNode{key: key, data: data}
	if l.tail == nil {
		l.tail = newNode
		l.head = newNode
		newNode.next = nil
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
		newNode.next = nil
	}
	l.size++
}

func (l *LinkedList) GetHead() *LinkedListNode {
	return l.head
}

func (l *LinkedList) GetTail() *LinkedListNode {
	return l.tail
}

func (l *LinkedList) RemoveNode(node *LinkedListNode) *LinkedListNode {
	if node == nil {
		return nil
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == l.head {

		l.head = l.head.next
	}
	if node == l.tail {
		l.tail = l.tail.prev
	}
	l.size--
	return node
}

func (l *LinkedList) Remove(data int) {
	i := l.GetHead()
	for i != nil {
		if i.data == data {
			l.RemoveNode(i)
		}
		i = i.next
	}
}

func (l *LinkedList) RemoveHead() *LinkedListNode {
	return l.RemoveNode(l.head)
}

func (l *LinkedList) RemoveTail() *LinkedListNode {
	return l.RemoveNode(l.tail)
}
