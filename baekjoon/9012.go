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

/*

func IsVPS(ps string) bool {
	stack := 0
	for _, char := range ps {
		if char == '(' {
			stack += 1
		} else {
			stack -= 1
			if stack == -1 { break }
		}
	}
	
	if stack == 0 {
		return true
	} else {
		return false
	}
}

*/
