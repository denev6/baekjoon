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

	var x, y int
	xPoints := make([]int, 2)
	yPoints := make([]int, 2)

	for i := 0; i < 3; i++ {
		fmt.Fscan(reader, &x, &y)
		xPoints = CheckPoint(x, xPoints)
		yPoints = CheckPoint(y, yPoints)
	}
	x = GetPoint(xPoints)
	y = GetPoint(yPoints)

	fmt.Fprint(writer, x, y)

}

func CheckPoint(newPoint int, arrPoints []int) []int {
	if newPoint == arrPoints[0] {
		arrPoints[0] = -1
	} else if newPoint == arrPoints[1] {
		arrPoints[1] = -1
	} else if arrPoints[0] == 0 {
		arrPoints[0] = newPoint
	} else {
		arrPoints[1] = newPoint
	}
	return arrPoints
}

func GetPoint(arrPoints []int) int {
	if arrPoints[0] != -1 {
		return arrPoints[0]
	} else {
		return arrPoints[1]
	}
}
