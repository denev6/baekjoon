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

	var a, b int
	for {
		_, err := fmt.Fscan(reader, &a, &b)
		if err != nil {
			break
		}
		fmt.Fprintln(writer, a+b)
	}
}
