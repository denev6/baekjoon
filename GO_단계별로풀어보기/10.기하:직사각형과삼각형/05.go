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

	var numPoint, x, y int
	fmt.Fscan(reader, &numPoint)

	xPoints := make([]int, numPoint)
	yPoints := make([]int, numPoint)

	for i := 0; i < numPoint; i++ {
		fmt.Fscan(reader, &x, &y)
		xPoints[i] = x
		yPoints[i] = y
	}

	xMin, xMax := GetMinMax(xPoints)
	yMin, yMax := GetMinMax(yPoints)
	area := (xMax - xMin) * (yMax - yMin)

	fmt.Fprint(writer, area)
}

func GetMinMax(arrPoints []int) (int, int) {
	// 문제 조건
	min := 100000
	max := -10000

	for _, point := range arrPoints {
		if point < min {
			min = point
		}
		if max < point {
			max = point
		}
	}
	return min, max
}
