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

	var n, m, start, end, num int
	fmt.Fscan(reader, &n, &m)
	
	arr := make([]int, n)

	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &start, &end, &num)
		for j := start - 1; j < end; j++ {
			arr[j] = num
		}
	}

	for _, res := range arr {
		fmt.Fprintf(writer, "%d ", res)
	}
}
