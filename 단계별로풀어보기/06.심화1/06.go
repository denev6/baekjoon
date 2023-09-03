package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)
var croatiaRegex = regexp.MustCompile(`c=|c-|dz=|d-|lj|nj|s=|z=`)

func main() {
	defer writer.Flush()

	var input string
	fmt.Fscan(reader, &input)

	input = croatiaRegex.ReplaceAllString(input, "$")
	fmt.Fprint(writer, len(input))
}
