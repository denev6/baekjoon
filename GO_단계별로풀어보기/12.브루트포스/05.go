package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n, cnt int
	fmt.Fscan(reader, &n)

	i := 666
	for ; cnt < n; i++ {
		if strings.Contains(strconv.Itoa(i), "666") {
			cnt += 1
		}
	}
	fmt.Fprint(writer, i-1)
}
