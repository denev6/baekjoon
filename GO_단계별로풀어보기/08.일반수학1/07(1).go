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

	var upHeight, downHeight, target int
	fmt.Fscan(reader, &upHeight, &downHeight, &target)

	target -= upHeight
	day := target/(upHeight-downHeight) + 1
	if (target % (upHeight - downHeight)) != 0 {
		day += 1
	}
	fmt.Fprint(writer, day)
}
