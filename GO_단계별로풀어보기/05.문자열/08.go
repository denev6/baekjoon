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

	sentence, _ := reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)
	
	if sentence != "" {
		words := strings.Split(sentence, " ")
		fmt.Fprint(writer, len(words))
	} else {
		fmt.Fprint(writer, 0)
	}
}
