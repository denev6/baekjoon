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

	var tmp int
	arr := []int{}

	for {
		_, err := fmt.Fscan(reader, &tmp)
		if err != nil {
			break
		}
		arr = append(arr, tmp)
	}

	// 문제 조건
	max := 0
	max_idx := -1

	for idx, n := range arr {
		if max < n {
			max = n
			max_idx = idx + 1
		}
	}
	fmt.Fprintln(writer, max)
	fmt.Fprint(writer, max_idx)
}
