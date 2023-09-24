package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	var numCard int 
	fmt.Fscan(reader, &numCard)
	
	deque := make([]int, numCard)
	for i := 0; i < numCard; i++ {
		deque[i] = i + 1
	}
	for len(deque) > 1 {
		deque = deque[1:]
		deque = append(deque, deque[0])
		deque = deque[1:]
	}
	fmt.Print(deque[0])
}
