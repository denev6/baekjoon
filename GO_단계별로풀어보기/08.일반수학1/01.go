package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var num string
	var base, res int
	fmt.Fscan(reader, &num, &base)

	for i, currentBase := len(num)-1, 1; i >= 0; i, currentBase = i-1, currentBase*base {
		res += CharToDecimal(rune(num[i])) * currentBase
	}
	fmt.Fprint(writer, res)
}

func CharToDecimal(char rune) int {
	if '0' <= char && char <= '9' {
		res, _ := strconv.Atoi(string(char))
		return res
	} else {
		return int(char-'A') + 10
	}
}
