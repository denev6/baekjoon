// 904KB, 12ms

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
	// 'A' = ASCII 65, 'A' ~ 'C' -> 2번 다이얼
	// 다이얼2 = (int('A') - int('A')) / 3 + 2
	// but, P~S, W~Z는 다이얼에 문자가 4개
	
	if 'Z' <= char {
		char -= 1
	}
	if 'S' <= char {
		char -= 1
	}
	index := int(char-'A')/3 + 2
	return index
}
