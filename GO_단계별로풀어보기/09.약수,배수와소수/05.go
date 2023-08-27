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

	var minRange, maxRange int
	fmt.Fscanln(reader, &minRange)
	fmt.Fscan(reader, &maxRange)

	allPrimes := GetPrimeNum(maxRange)[minRange-1:]
	sumPrimes := 0
	minPrime := 0

	for idx, isPrime := range allPrimes {
		if isPrime {
			sumPrimes += idx + minRange
			if minPrime == 0 {
				minPrime = idx + minRange
			}
		}
	}

	if minPrime == 0 {
		fmt.Fprint(writer, -1)
	} else {
		fmt.Fprintf(writer, "%d\n%d", sumPrimes, minPrime)
	}
}

func GetPrimeNum(maxRange int) []bool {
	// 에라토스테네스의 체
	isPrime := make([]bool, maxRange)
	for idx, _ := range isPrime {
		isPrime[idx] = true
	}
	isPrime[0] = false  // 1은 소수가 아님

	sqrt := int(math.Sqrt(float64(maxRange)))
	for i := 2; i <= sqrt; i++ {
		if isPrime[i-1] {
			for j := 2; j*i <= maxRange; j++ {
				isPrime[i*j-1] = false
			}
		}
	}
	return isPrime
}
