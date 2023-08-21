package main

import (
	"fmt"
)

func main() {
	var numCases, a, b int
	fmt.Scan(&numCases)
	for i := 1; i < numCases+1; i++ {
		fmt.Scan(&a, &b)
		fmt.Printf("Case #%d: %d + %d = %d\n", i, a, b, a+b)
	}
}
