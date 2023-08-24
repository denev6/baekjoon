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

	var n, m, start, end int
	fmt.Fscan(reader, &n, &m)
	baskets := InitArr(n)

	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &start, &end)
		start -= 1
		baskets = ReversePartialSlice(baskets, start, end)
	}
	PrintSlice(baskets)
}

func InitArr(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	return arr
}

func ReversePartialSlice(intSlice []int, start, end int) []int {
	sliced := intSlice[start:end]
	reversed := ReverseSlice(sliced)
	for i, j := start, 0; i < end; i, j = i+1, j+1 {
		intSlice[i] = reversed[j]
	}
	return intSlice
}

func ReverseSlice(intSlice []int) []int {
	for i, j := 0, len(intSlice)-1; i < j; i, j = i+1, j-1 {
		intSlice[i], intSlice[j] = intSlice[j], intSlice[i]
	}
	return intSlice
}

func PrintSlice(intSlice []int) {
	for _, n := range intSlice {
		fmt.Fprintf(writer, "%d ", n)
	}
}
