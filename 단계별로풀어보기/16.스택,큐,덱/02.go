package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	var numInput, input int
	fmt.Fscan(reader, &numInput)

	stack := make([]int, numInput)
	lenStack := 0
	sum := 0

	for i := 0; i < numInput; i++ {
		fmt.Fscan(reader, &input)
		if input == 0 {
			lenStack--
			sum -= stack[lenStack]
		} else {
			stack[lenStack] = input
			sum += input
			lenStack++
		}
	}
	fmt.Print(sum)
}
