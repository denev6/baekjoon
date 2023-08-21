package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var numCases, a, b int
	fmt.Fscan(reader, &numCases)
	for i := 0; i < numCases; i++ {
		fmt.Fscan(reader, &a, &b)
		fmt.Fprintln(writer, a+b)
	}
}
