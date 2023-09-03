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

	var step int 
	fmt.Fscan(reader, &step)

	dotPerSide := 1
	for i := 0; i < step; i++ {
		dotPerSide *= 2
	}
	numDot := (dotPerSide + 1) * (dotPerSide + 1)
	fmt.Fprint(writer, numDot)
}
