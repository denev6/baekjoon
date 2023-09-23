// container/list를 이용한 추상화 풀이
// 시간초과

package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	stack := NewStack()

	sc.Scan()
	numCommand, _ := strconv.Atoi(sc.Text())

	for i := 0; i < numCommand; i++ {
		sc.Scan()
		command := strings.Trim(sc.Text(), " ")

		if command[0] == '1' {
			valToPush := strings.Split(command, " ")[1]
			numToPush, _ := strconv.Atoi(valToPush)
			stack.Push(numToPush)
			continue
		}

		var res interface{}
		if command == "2" {
			res = stack.Pop(true)
		} else if command == "3" {
			res = stack.Len()
		} else if command == "4" {
			res = stack.IsEmpty()
		} else if command == "5" {
			res = stack.Pop(false)
		}
		fmt.Println(res)
	}
}

type Stack struct {
	List *list.List
}

func NewStack() *Stack {
	stack := Stack{list.New()}
	return &stack
}

func (q *Stack) Len() int {
	return q.List.Len()
}

func (q *Stack) IsEmpty() int {
	if q.List.Len() == 0 {
		return 1
	}
	return 0

}

func (q *Stack) Push(n int) {
	q.List.PushBack(n)
}

func (q *Stack) Pop(toRemove bool) interface{} {
	val := q.List.Back() // val: *List.Element
	if val != nil {
		if toRemove {
			return q.List.Remove(val) // List.Element.Value
		}
		return val.Value
	}
	return -1
}
