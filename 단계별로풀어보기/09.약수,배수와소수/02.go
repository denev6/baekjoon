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

	var n, k, res int
	fmt.Fscan(reader, &n, &k)

	divisor := 1
	cnt := 0
	
	for ; divisor <= n && cnt < k; divisor++ {
		if n%divisor == 0 {
			cnt++
		}
	}
	if cnt < k {
	  res = 0
	} else {
	  res = divisor - 1
	}
	fmt.Fprint(writer, res)
}
