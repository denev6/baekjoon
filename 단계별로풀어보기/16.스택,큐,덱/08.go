package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var numElement, targetIdx int
	fmt.Fscan(reader, &numElement, &targetIdx)
	targetIdx -= 1

	list := NewCircularList(numElement)
	toPrint := make([]int, 0, numElement)

	// Use "Circular Linked List"
	for !list.IsEmpty() {
		for i := 0; i < targetIdx; i++ {
			list.Next()
		}
		removed := list.RemoveCurrent()
		toPrint = append(toPrint, removed)
	}

	// Formatting result
	fmt.Fprintf(writer, "<%d", toPrint[0])
	for i := 1; i < numElement; i++ {
		fmt.Fprintf(writer, ", %d", toPrint[i])
	}
	fmt.Fprint(writer, ">")
}

// Circular Linked List
type CircularList struct {
	list    []int
	pointer int
	len     int
}

// Init CircularList with ordered array.
// i.e. 1-2-3-4- ... -1
func NewCircularList(length int) *CircularList {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = i + 1
	}
	len_arr := len(arr)
	list := CircularList{arr, 0, len_arr}
	return &list
}

func (l *CircularList) IsEmpty() bool {
	if l.len == 0 {
		return true
	}
	return false
}

func (l *CircularList) Next() {
	if l.pointer == l.len-1 {
		l.pointer = 0
	} else {
		l.pointer++
	}
}

// Remove current element
// and move pointer to next element.
func (l *CircularList) RemoveCurrent() int {
	target := l.list[l.pointer]
	l.list = append(l.list[:l.pointer], l.list[l.pointer+1:]...)
	l.len -= 1
	if l.pointer > l.len-1 {
		l.pointer = 0
	}
	return target
}
