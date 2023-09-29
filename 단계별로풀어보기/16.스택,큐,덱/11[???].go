// 왜 틀렸는지 알 수 없는 풀이...
// 예제는 모두 통과 

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Scan() // Skip
	structures := strings.Split(ReadLine(), " ")
	items := strings.Split(ReadLine(), " ")
	sc.Scan() // Skip
	newItems := strings.Split(ReadLine(), " ")

	qs := NewQueueStack(structures, items)
	res := make([]string, len(newItems))
	for i, new := range newItems {
		res[i] = qs.Insert(new)
	}
	fmt.Print(strings.Join(res, " "))
}

func ReadLine() string {
	sc.Scan()
	return sc.Text()
}

// Just Queue filtering Stack values
type QueueStack struct {
	queue []string
	len   int
}

func NewQueueStack(flag, items []string) *QueueStack {
	validItems := make([]string, 0, len(flag))
	for i, f := range flag {
		if f == "0" {
			validItems = append(validItems, items[i])
		}
	}
	return &QueueStack{validItems, len(validItems)}
}

func (s *QueueStack) Insert(new string) string {
	if s.len == 0 {
		return new 
	}
	res := s.queue[s.len-1]
	s.appendLeft(new)
	return res 
}

func (s *QueueStack) appendLeft(new string) {
	s.queue = append(s.queue, " ")
	copy(s.queue[1:], s.queue)
	s.queue[0] = new 
}