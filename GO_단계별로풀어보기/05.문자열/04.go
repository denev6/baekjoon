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

	var char rune // ASCII
	// fmt.Fscan(reader, &char) -> Doesn't work
	fmt.Fscanf(reader, "%c", &char)
	fmt.Fprint(writer, char)
}
