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

	var n, m, a, b int
	fmt.Fscan(reader, &n, &m)
	
	arr := InitArr(n)

	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a, &b)
		a -= 1
		b -= 1
		arr[a], arr[b] = arr[b], arr[a]
	}

	for _, res := range arr {
		fmt.Fprintf(writer, "%d ", res)
	}
}

func InitArr(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	return arr
}
