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

	var step, upper, lower int
	fmt.Fscan(reader, &step)

	level := 1
	for ; level < step; level++ {
		step -= level
	}

	if level%2 == 0 {
		upper = step
		lower = level - step + 1
	} else {
		upper = level - step + 1
		lower = step
	}
	fmt.Fprintf(writer, "%d/%d", upper, lower)
}
