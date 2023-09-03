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

	var num int

	maxNum := -1
	maxNumColIndex := -1
	maxNumRowIndex := -1

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Fscan(reader, &num)
			if maxNum < num {
				maxNum = num
				maxNumColIndex = i
				maxNumRowIndex = j
			}
		}
	}
	fmt.Fprintln(writer, maxNum)
	fmt.Fprint(writer, maxNumColIndex+1, maxNumRowIndex+1)
}
