package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var numPoint, x, y int
	fmt.Fscan(reader, &numPoint)

	arr := make([]point, numPoint)

	for i := 0; i < numPoint; i++ {
		fmt.Fscan(reader, &x, &y)
		arr[i] = point{x, y}
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i].y == arr[j].y {
			return arr[i].x < arr[j].x
		}
		return arr[i].y < arr[j].y
	})

	for _, p := range arr {
		fmt.Fprintln(writer, p.x, p.y)
	}
}

type point struct {
	x int
	y int
}
