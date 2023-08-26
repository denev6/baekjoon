package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var input int
	fmt.Fscan(reader, &input)
	PrintStars(input)
}

func PrintStars(level int) {
	stars := make([]string, level)
	for i := 1; i <= level; i++ {
		line := ""
		line += strings.Repeat(" ", level-i)
		line += strings.Repeat("*", i*2-1)
		fmt.Fprintln(writer, line)
		stars[i-1] = line
	}
	for j := level - 1; j > 0; j-- {
		fmt.Fprintln(writer, stars[j-1])
	}
}
