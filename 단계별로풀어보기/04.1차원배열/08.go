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

	var input, remainder int
	var remainders []int

	for {
		_, err := fmt.Fscan(reader, &input)
		if err != nil {
			break
		}
		remainder = input % 42
		remainders = append(remainders, remainder)
	}
	res := removeDuplicateInt(remainders)
	fmt.Fprint(writer, len(res))
}

func removeDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	slice := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			slice = append(slice, item)
		}
	}
	return slice
}
