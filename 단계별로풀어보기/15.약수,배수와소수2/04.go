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

	var numTree, point, previousPoint int
	fmt.Fscan(reader, &numTree)
	numTree -= 1

	points := make([]int, numTree)
	fmt.Fscan(reader, &previousPoint)

	fisrtPoint := previousPoint
	for i := 0; i < numTree; i++ {
		fmt.Fscan(reader, &point)
		points[i] = point - previousPoint
		previousPoint = point
	}
	lastPoint := point

	gcd := points[0]
	for _, p := range points[1:] {
		for p > 0 {
			gcd, p = p, gcd%p
		}
	}

	numAllTree := (lastPoint - fisrtPoint) / gcd
	numToPlant := numAllTree - numTree

	fmt.Fprintln(writer, numToPlant)
}
