package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var lenArr, tmp int
	fmt.Fscan(reader, &lenArr)

	arr := make([]int, lenArr)

	for i := 0; i < lenArr; i++ {
		fmt.Fscan(reader, &tmp)
		arr[i] = tmp
	}

	// Golang Default Sort
	sort.Ints(sort.IntSlice(arr))

	for _, n := range arr {
		fmt.Fprintf(writer, "%d\n", n)
	}
}
