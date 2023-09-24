// 시간 단축을 위한 풀이
// 통과

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

	var numCommand, command, num int
	fmt.Fscan(reader, &numCommand)

	deque := NewDeque()

	for i := 0; i < numCommand; i++ {
		fmt.Fscan(reader, &command)
		switch command {
		case 1:
			fmt.Fscan(reader, &num)
			deque.PushFront(num)
		case 2:
			fmt.Fscan(reader, &num)
			deque.PushBack(num)
		case 3:
			num = deque.PopFront()
		case 4:
			num = deque.PopBack()
		case 5:
			num = deque.Len()
		case 6:
			num = deque.IsEmpty()
		case 7:
			num = deque.Head()
		case 8:
			num = deque.Tail()
		}
		if (command == 1) || (command == 2) {
			continue
		}
		fmt.Fprintln(writer, num)
	}
}

// Deque implemented in Go-Slice
type Deque struct {
	list []int
	len  int
}

func NewDeque() *Deque {
	deque := Deque{[]int{}, 0}
	return &deque
}

func (d *Deque) PushFront(n int) {
	d.list = append(d.list, 0)
	copy(d.list[1:], d.list)
	d.list[0] = n
	d.len += 1
}

func (d *Deque) PushBack(n int) {
	d.list = append(d.list, n)
	d.len += 1
}

func (d *Deque) Len() int {
	return d.len
}

func (d *Deque) IsEmpty() int {
	if d.len == 0 {
		return 1
	} else {
		return 0
	}
}

func (d *Deque) PopFront() int {
	if d.len == 0 {
		return -1
	}
	target := d.list[0]
	d.list = d.list[1:]
	d.len--
	return target
}

func (d *Deque) Head() int {
	if d.len == 0 {
		return -1
	}
	return d.list[0]
}

func (d *Deque) PopBack() int {
	if d.len == 0 {
		return -1
	}
	target := d.list[d.len-1]
	d.list = d.list[:d.len-1]
	d.len--
	return target
}

func (d *Deque) Tail() int {
	if d.len == 0 {
		return -1
	}
	return d.list[d.len-1]
}
