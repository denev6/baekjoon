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

	var numCard, target, card, sum int
	fmt.Fscanln(reader, &numCard, &target)

	cards := make([]int, numCard)
	for i := 0; i < numCard; i++ {
		fmt.Fscan(reader, &card)
		cards[i] = card
	}

	maxRes := -1
	for i := 0; i < numCard-2; i++ {
		for j := i + 1; j < numCard-1; j++ {
			for k := j + 1; k < numCard; k++ {
				sum = cards[i] + cards[j] + cards[k]
				if sum <= target && maxRes < sum {
					maxRes = sum
				}
			}
		}
	}

	fmt.Fprint(writer, maxRes)
}
