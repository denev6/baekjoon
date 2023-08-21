package main

import (
	"fmt"
)

func main() {
	var n int
	var line string
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		line = RepeatChar(n-i, " ") + RepeatChar(i, "*")
		fmt.Println(line)
	}
}

func RepeatChar(n int, char string) string {
	res := ""
	for i := 0; i < n; i++ {
		res += char
	}
	return res
}
