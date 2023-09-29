package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var b strings.Builder

func main() {
	sc.Scan()
	lenList, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	items := strings.Split(sc.Text(), " ")
	list := NewCircularList(lenList)

	var idx int
	for i := 0; i < lenList-1; i++ {
		list.DeleteCurrent()
		idx := list.Index()

		b.WriteString(strconv.Itoa(idx + 1))
		b.WriteString(" ")

		item, _ := strconv.Atoi(items[idx])
		for item != 0 {
			if item > 0 {
				list.Next()
				item--
			} else {
				list.Previous()
				item++
			}
		}
	}
	idx = list.Index()
	b.WriteString(strconv.Itoa(idx + 1))
	fmt.Print(b.String())
}

// Circular Linked List
type CircularList struct {
	isDeleted []bool
	pointer   int
	len       int
}

func NewCircularList(length int) *CircularList {
	arr := make([]bool, length)
	return &CircularList{arr, 0, length}
}

func (l *CircularList) Index() int {
	return l.pointer
}

func (l *CircularList) Previous() {
	if l.pointer == 0 {
		l.pointer = l.len - 1
	} else {
		l.pointer--
	}
	for l.isDeleted[l.pointer] == true {
		if l.pointer == 0 {
			l.pointer = l.len - 1
		} else {
			l.pointer--
		}
	}
}

func (l *CircularList) Next() {
	if l.pointer == l.len-1 {
		l.pointer = 0
	} else {
		l.pointer++
	}
	for l.isDeleted[l.pointer] == true {
		if l.pointer == l.len-1 {
			l.pointer = 0
		} else {
			l.pointer++
		}
	}
}

// Delete current item
func (l *CircularList) DeleteCurrent() {
	l.isDeleted[l.pointer] = true
}
