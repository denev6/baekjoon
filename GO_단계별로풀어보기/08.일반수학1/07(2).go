// 시간초과

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

	var upHeight, downHeight, target, currentHeight int
	fmt.Fscan(reader, &upHeight, &downHeight, &target)

	day := 1
	for ; ; day += 1 {
		currentHeight += upHeight
		if target <= currentHeight {
			break
		}
		currentHeight -= downHeight
	}
	fmt.Fprint(writer, day)
}
