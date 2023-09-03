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

	var numInput, numIter int
	var word string

	fmt.Fscanln(reader, &numInput)
	for i := 0; i < numInput; i++ {
		fmt.Fscanln(reader, &numIter, &word)
		IterChar(numIter, word)
	}
}

func IterChar(numIter int, word string) {
	var res string
	for _, char := range word {
		for i := 0; i < numIter; i++ {
			res += string(char)
		}
	}
	fmt.Fprintln(writer, res)
}
