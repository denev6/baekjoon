package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	var n int
	cases := make([]int, 0, 10)

	minN := 123457
	maxN := 0
	for {
		fmt.Fscan(reader, &n)
		if n == 0 {
			break
		}

		cases = append(cases, n)
		if n < minN {
			minN = n
		}
		if maxN < n {
			maxN = n
		}
	}

	allPrimes := GetPrimeNum(minN, maxN*2)
	for _, n := range cases {
		cnt := 0
		for _, isPrime := range allPrimes[n-minN+1 : n*2-minN+1] {
			if isPrime {
				cnt += 1
			}
		}
		fmt.Println(cnt)
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
