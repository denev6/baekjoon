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

	var total int
	fmt.Fscan(reader, &total)

	fiveKg := total / 5
	left := total % 5
	threeKg := 0

	for ; 0 <= fiveKg; left, fiveKg = left+5, fiveKg-1 {
		if left%3 == 0 {
			threeKg = left / 3
			break
		}
	}
	if fiveKg*5+threeKg*3 == total {
		fmt.Fprint(writer, fiveKg+threeKg)
	} else {
		fmt.Fprint(writer, -1)
	}
}
