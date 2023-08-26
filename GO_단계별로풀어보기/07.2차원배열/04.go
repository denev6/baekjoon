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

	var numPatch, row, col, area int
	board := make([][100]bool, 100)

	fmt.Fscan(reader, &numPatch)
	for i := 0; i < numPatch; i++ {
		fmt.Fscan(reader, &row, &col)
		for c := col - 1; c < col+9; c++ {
			for r := row - 1; r < row+9; r++ {
				board[c][r] = true
			}
		}
	}

	for _, boardRow := range board {
		for _, isCovered := range boardRow {
			if isCovered {
				area += 1
			}
		}
	}

	fmt.Fprint(writer, area)
}
