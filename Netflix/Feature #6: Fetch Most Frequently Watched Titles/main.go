package main

import (
	"fmt"
)

type LFUStructure struct {
	capacity int
	size int
	minFreq int
	freqDict map[int]*LinkedList
	keyDict map[int]*LinkedListNode       
}

func (lf *LFUStructure) Get(key int) *LinkedListNode{
	if _, ok := lf.keyDict[key]; !ok {
		return nil
	}
	temp := lf.keyDict[key]

	lf.freqDict[temp.freq].DeleteNode(temp)
	if lf.freqDict[lf.keyDict[key].freq].head == nil {
		delete(lf.freqDict, lf.keyDict[key].freq)
		if lf.minFreq == lf.keyDict[key].freq {
			lf.minFreq++
		}
	}
	lf.keyDict[key].freq++

	if _, ok := lf.freqDict[lf.keyDict[key].freq]; !ok {
		lf.freqDict[lf.keyDict[key].freq] = &LinkedList{}
	}
	lf.freqDict[lf.keyDict[key].freq].Append(lf.keyDict[key])
	return lf.keyDict[key]
}

func (lf *LFUStructure) Set(key, value int){
	if lf.Get(key) != nil {
		lf.keyDict[key].val = value
		return
	}
	if lf.size == lf.capacity {
		delete(lf.keyDict, lf.freqDict[lf.minFreq].head.key)
		lf.freqDict[lf.minFreq].DeleteNode(lf.freqDict[lf.minFreq].head)
		if lf.freqDict[lf.minFreq].head == nil {
			delete(lf.freqDict, lf.minFreq)
		}
		lf.size--
	}
	lf.minFreq = 1
	lf.keyDict[key] = &LinkedListNode{key: key, val: value, freq: lf.minFreq}
	if _, ok := lf.freqDict[lf.keyDict[key].freq]; !ok {
		lf.freqDict[lf.keyDict[key].freq] = &LinkedList{}
	}
	lf.freqDict[lf.keyDict[key].freq].Append(lf.keyDict[key])
	lf.size++
}

func (lf *LFUStructure) PrintDict(){
	for first, second := range lf.keyDict {
		fmt.Printf("(%v,%v)", first, second.val)
	}
	fmt.Println("")
}

func main() {
  obj := &LFUStructure{capacity: 2, freqDict : make(map[int]*LinkedList), keyDict : make(map[int]*LinkedListNode)}
  fmt.Println("The most frequently watched titles are: (key, value)")
  obj.Set(1, 1);
  obj.Set(2, 2);
  obj.PrintDict();
  obj.Get(1);
  obj.Set(3, 3);
  obj.PrintDict();
  obj.Get(2);
  obj.Set(4, 4);

  obj.Get(1);
  obj.Get(3);
  obj.Get(4);
  obj.PrintDict();
}