package main

import (
	"fmt"
)

type Node struct {
	Data int
	next *Node
}

type SingleLinkedList struct {
	root  *Node
	tail  *Node
	count int
}

func NewLinkedList() *SingleLinkedList {
	return &SingleLinkedList{count: 0}
}

func (linkedList *SingleLinkedList) Len() int {
	return linkedList.count
}

func (linkedList *SingleLinkedList) Append(value int) {
	node := &Node{Data: value}
	if linkedList.root == nil {
		linkedList.root = node
		linkedList.tail = node
		linkedList.count = 1
		return
	}
	linkedList.tail.next = node
	linkedList.tail = node
	linkedList.count += 1
}

func (linkedList *SingleLinkedList) IsIncluded(data int) bool {
	currentNode := linkedList.root
	for ; currentNode != nil; currentNode = currentNode.next {
		if currentNode.Data == data {
			return true
		}
	}
	return false
}

func main() {
	dataArray := []int{3, 6, 1, 2, 9, 10}
	linkedList := NewLinkedList()

	// Append
	for _, num := range dataArray {
		linkedList.Append(num)
	}

	// Length
	fmt.Println("Length:", linkedList.Len())

	// Search
	isExist := linkedList.IsIncluded(21)
	if isExist {
		fmt.Println("Found 21!")
	}
	isExist = linkedList.IsIncluded(9)
	if isExist {
		fmt.Println("Found 9!")
	}

}
