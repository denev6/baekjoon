package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var numCase, n int
	fmt.Fscan(reader, &numCase)

	for i := 0; i < numCase; i++ {
		fmt.Fscan(reader, &n)

		if n == 0 || n == 1 {
			fmt.Fprintln(writer, 2)
			continue
		}

		for ; ; n++ {
			if isPrime(n) {
				break
			}
		}
		fmt.Fprintln(writer, n)
	}
}

func isPrime(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	for div := 2; div <= sqrt; div++ {
		if n%div == 0 {
			return false
		}
	}
	return true
}
