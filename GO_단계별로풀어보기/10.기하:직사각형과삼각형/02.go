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
	var x, y, w, h, minDist int
	fmt.Fscan(reader, &x, &y, &w, &h)

	distHorizontal := MinDist(x, w)
	distVertical := MinDist(y, h)

	if distHorizontal < distVertical {
		minDist = distHorizontal
	} else {
		minDist = distVertical
	}
	fmt.Fprint(writer, minDist)
}

func MinDist(current, maxPoint int) int {
	distToMax := maxPoint - current
	if distToMax < current {
		return distToMax
	} else {
		return current
	}
}
