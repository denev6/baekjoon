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

	var str string
	fmt.Fscan(reader, &str)

	patterns := make(map[string]struct{})

	for i := 0; i < len(str); i++ {
		var tmp string
		for _, char := range str[i:] {
			tmp += string(char)
			patterns[tmp] = struct{}{}
		}
	}
	fmt.Fprint(writer, len(patterns))
}
