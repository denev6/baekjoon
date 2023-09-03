// 900KB, 4ms

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

	var input string
	fmt.Fscan(reader, &input)

	res := 0
	defaultSec := 1
	for _, char := range input {
		res += CharToDial(char) + defaultSec
	}
	fmt.Fprint(writer, res)
}

func CharToDial(char rune) int {
	var dial int
	if 'A' <= char && char <= 'C' {
		dial = 2
	} else if char <= 'F' {
		dial = 3
	} else if char <= 'I' {
		dial = 4
	} else if char <= 'L' {
		dial = 5
	} else if char <= 'O' {
		dial = 6
	} else if char <= 'S' {
		dial = 7
	} else if char <= 'V' {
		dial = 8
	} else {
		dial = 9
	}
	return dial
}
