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

	var input int
	answer := [...]int{1, 1, 2, 2, 2, 8}
	res := [6]int{}

	for i, n := range answer {
		fmt.Fscan(reader, &input)
		res[i] = n - input
	}
	PrintArr(res)
}

func PrintArr(arr [6]int) {
	for _, n := range arr {
		fmt.Fprintf(writer, "%d ", n)
	}
}
