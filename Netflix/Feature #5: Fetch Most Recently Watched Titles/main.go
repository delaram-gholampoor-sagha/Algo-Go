package main

import (
	"fmt"
)

type LRUCache struct {
	capacity int

	//LinkedListNode holds key and value pairs
	cache     map[int]*LinkedListNode
	cacheVals LinkedList
}

func (lr *LRUCache) Get(key int) *LinkedListNode {
	if _, ok := lr.cache[key]; !ok {
		return nil
	} else {
		value := lr.cache[key].data
		lr.cacheVals.RemoveNode(lr.cache[key])
		lr.cacheVals.InsertAtTail(key, value)
		return lr.cacheVals.GetTail()
	}
}

func (lr *LRUCache) Set(key, value int) {
	if _, ok := lr.cache[key]; !ok {
		if lr.cacheVals.size >= lr.capacity {
			lr.cacheVals.InsertAtTail(key, value)
			lr.cache[key] = lr.cacheVals.GetTail()
			delete(lr.cache, lr.cacheVals.GetHead().key)
			lr.cacheVals.RemoveHead()
		} else {
			lr.cacheVals.InsertAtTail(key, value)
			lr.cache[key] = lr.cacheVals.GetTail()
		}
	} else {
		lr.cacheVals.RemoveNode(lr.cache[key])
		lr.cacheVals.InsertAtTail(key, value)
		lr.cache[key] = lr.cacheVals.GetTail()
	}
}

func (lr *LRUCache) Print() {
	curr := lr.cacheVals.head
	for curr != nil {
		fmt.Printf("(%v,%v)", curr.key, curr.data)
		curr = curr.next
	}
	fmt.Println("")
}

func main() {
	cache := &LRUCache{capacity: 3, cache: make(map[int]*LinkedListNode)}
	fmt.Println("The most recently watched titles are: (key, value)")
	cache.Set(10, 20)
	cache.Print()

	cache.Set(15, 25)
	cache.Print()

	cache.Set(20, 30)
	cache.Print()

	cache.Set(25, 35)
	cache.Print()

	cache.Set(5, 40)
	cache.Print()

	cache.Get(25)
	cache.Print()
}
