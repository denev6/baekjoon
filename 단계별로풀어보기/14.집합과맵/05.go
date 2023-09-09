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

	var numCard, numInput int
	var tmp int
	
	cards := make(map[int]int)
	
	fmt.Fscan(reader, &numCard)
	for i := 0; i < numCard; i++ {
		fmt.Fscan(reader, &tmp)

		if _, exists := cards[tmp]; exists {
			cards[tmp] += 1
		} else {
			cards[tmp] = 1
		}
	}

	fmt.Fscan(reader, &numInput)
	for i := 0; i < numInput; i++ {
		fmt.Fscan(reader, &tmp)

		res := 0
		if _, includes := cards[tmp]; includes {
			res = cards[tmp]
		}
		fmt.Fprintf(writer, "%d ", res)
	}
}
