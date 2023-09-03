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

	var input string 
	fmt.Fscan(reader, &input)

	arr := []rune(input)

	// Golang Default Sort
	// https://pkg.go.dev/sort#Slice
	sort.Slice(arr, func(i, j int) bool {
        return arr[i] > arr[j]
    })

	res := string(arr)
	fmt.Fprint(writer, res)
}
