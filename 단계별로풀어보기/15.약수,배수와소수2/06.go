package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	var minRange, maxRange int
	fmt.Fscan(reader, &minRange, &maxRange)

	allPrimes := GetPrimeNum(minRange, maxRange)
	for idx, isPrime := range allPrimes {
		if isPrime {
			fmt.Println(idx + minRange)
		}
	}
}

func GetPrimeNum(minRange, maxRange int) []bool {
	isPrime := make([]bool, maxRange)
	for idx, _ := range isPrime {
		isPrime[idx] = true
	}
	isPrime[0] = false

	sqrt := int(math.Sqrt(float64(maxRange)))
	for i := 2; i <= sqrt; i++ {
		if isPrime[i-1] {
			for j := 2; j*i <= maxRange; j++ {
				isPrime[i*j-1] = false
			}
		}
	}
	return isPrime[minRange-1:]
}
