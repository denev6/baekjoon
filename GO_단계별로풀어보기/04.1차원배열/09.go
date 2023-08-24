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
		baskets = ReverseSlice(baskets, start, end)
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

func ReverseSlice(intSlice []int, start, end int) []int {
	sliced := intSlice[start:end]
	reversed := Reverse(sliced)
	for i, j := start, 0; i < end; i, j = i+1, j+1 {
		intSlice[i] = reversed[j]
	}
	return intSlice
}

func Reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func PrintSlice(intSlice []int) {
	for _, n := range intSlice {
		fmt.Fprintf(writer, "%d ", n)
	}
}
