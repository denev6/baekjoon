package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	var numPeople, order int
	fmt.Fscan(reader, &numPeople)

	stack := []int{}
	queue := make([]int, numPeople)
	for i := 0; i < numPeople; i++ {
		fmt.Fscan(reader, &order)
		queue[i] = order
	}

	currentOrder := 1
	for currentOrder <= numPeople {
		if (len(queue) > 0) && (currentOrder == queue[0]) {
			queue = queue[1:]
			currentOrder++
		} else if (len(stack) > 0) && (currentOrder == stack[len(stack)-1]) {
			stack = stack[:len(stack)-1]
			currentOrder++
		} else {
			if len(queue) == 0 {
				fmt.Print("Sad")
				return
			}
			stack = append(stack, queue[0])
			queue = queue[1:]
		}
	}
	fmt.Print("Nice")
}
