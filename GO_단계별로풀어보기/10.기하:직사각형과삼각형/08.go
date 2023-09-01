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
	
	var a, b, c int 
	fmt.Fscan(reader, &a, &b, &c)

	length := maxValidLegnth(a, b, c)
	fmt.Fprint(writer, length)
}

func maxValidLegnth(a, b, c int) int {
	max := -1
	sumAll := a + b + c

	for _, n := range []int{a, b, c} {
		if max < n {
			max = n
		}
	}

	if max < sumAll-max {
		return sumAll
	} else {
		return (sumAll-max) * 2 - 1
	}
}