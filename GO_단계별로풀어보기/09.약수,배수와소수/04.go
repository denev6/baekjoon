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

	var numCase, n, numPrime int
	fmt.Fscan(reader, &numCase)

	for i := 0; i < numCase; i++ {
		fmt.Fscan(reader, &n)
		if n == 1 {
			continue
		}
		isPrime := true

		sqrt := int(math.Sqrt(float64(n)))
		for div := 2; div <= sqrt; div++ {
			if n%div == 0 {
				isPrime = false
			}
		}
		if isPrime {
			numPrime += 1
		}
	}
	fmt.Fprint(writer, numPrime)
}
