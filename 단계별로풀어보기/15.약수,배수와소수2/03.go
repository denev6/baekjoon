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

	var a1, b1, a2, b2 int
	fmt.Fscan(reader, &a1, &b1)
	fmt.Fscan(reader, &a2, &b2)

	denominator := b1 * b2
	numerator := a1*b2 + a2*b1

	a, b := numerator, denominator
	for b > 0 {
		a, b = b, a%b
	}
	gcd := a

	fmt.Fprintln(writer, numerator/gcd, denominator/gcd)
}
