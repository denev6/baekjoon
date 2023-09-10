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

	var a, b int
	fmt.Fscan(reader, &a, &b)

	// 유클리드 호제법
	lcm := a * b
	for b > 0 {
		a, b = b, a%b
	}
	// a: 최대공약수
	lcm /= a
	fmt.Fprintln(writer, lcm)
}
