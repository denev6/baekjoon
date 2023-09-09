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

	var numSetA, numSetB, tmp int
	fmt.Fscan(reader, &numSetA, &numSetB)

	setA := make(map[int]struct{})

	for i := 0; i < numSetA; i++ {
		fmt.Fscan(reader, &tmp)
		setA[tmp] = struct{}{}
	}

	for i := 0; i < numSetB; i++ {
		fmt.Fscan(reader, &tmp)
		if _, exists := setA[tmp]; exists {
			numSetA -= 1
		} else {
			numSetA += 1
		}
	}

	fmt.Fprint(writer, numSetA)
}
