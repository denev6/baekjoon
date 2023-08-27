package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n, tmp int
	for {
		fmt.Fscan(reader, &n)
		if n == -1 {
			break
		}
		
		var res string 
		var factors []int
		sqrt := int(math.Sqrt(float64(n)))
		sumFactors := 0

		for divisor := 1; divisor <= sqrt; divisor++ {
			if n%divisor == 0 {
				sumFactors += divisor
				factors = append(factors, divisor)
			}
		}
		for i := len(factors); 0 < i; i-- {
			tmp = n / factors[i-1]
			sumFactors += tmp
			factors = append(factors, tmp)
		}

		if sumFactors == n * 2 {
			res = "= "
			numFactors := len(factors) - 2
			for idx := 0; idx < len(factors)-1; idx++ {
				res += strconv.Itoa(factors[idx])
				if idx != numFactors {
					res += " + "
				}
			}
		} else {
			res = "is NOT perfect."
		}
		fmt.Fprintln(writer, n, res)
	}
}
