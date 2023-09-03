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

	var input, res string
	var maxLen int

	inputs := make([]string, 5)

	for i := 0; i < 5; i++ {
		fmt.Fscan(reader, &input)
		inputs[i] = input
		if maxLen < len(input) {
			maxLen = len(input)
		}
	}

	for j := 0; j < maxLen; j++ {
		for _, str := range inputs {
			if j < len(str) {
				res += string(str[j])
			}
		}
	}

	fmt.Fprint(writer, res)
}
