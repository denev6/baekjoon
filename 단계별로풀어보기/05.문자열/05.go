package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var numInt, tmp int
	var ints string
	fmt.Fscan(reader, &numInt)
	fmt.Fscan(reader, &ints)

	sum := 0
	for _, n := range ints {
		tmp, _ = strconv.Atoi(string(n))
		sum += tmp
	}
	fmt.Fprint(writer, sum)
}
