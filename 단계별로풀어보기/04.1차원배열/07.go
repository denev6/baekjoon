package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	arr := [30]bool{}

	var num int
	for i := 0; i < 30; i++ {
		fmt.Fscan(reader, &num)
		arr[num-1] = true
	}

	for idx, isAttended := range arr {
		if !isAttended {
			fmt.Fprintln(writer, idx+1)
		}
	}
}
