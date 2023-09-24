// 4ms, 928KB

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// (V)PS: (Valid) Parenthesis String
	var numPs int
	var ps string

	fmt.Fscanln(reader, &numPs)

	for i := 0; i < numPs; i++ {
		fmt.Fscanln(reader, &ps)

		if IsVPS(ps) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func IsVPS(ps string) bool {
	for strings.Contains(ps, "()") {
		ps = strings.ReplaceAll(ps, "()", "")
	}
	if len(ps) == 0 {
		return true
	} else {
		return false
	}
}
