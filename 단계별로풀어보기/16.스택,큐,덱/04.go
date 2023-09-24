package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	for {
		sc.Scan()
		sentence := sc.Text()
		if sentence == "." {
			return
		}
		if IsVstr(sentence) {
			fmt.Fprintln(writer, "yes")
		} else {
			fmt.Fprintln(writer, "no")
		}
	}
}

func IsVstr(str string) bool {
	stack := []rune{}
	for _, char := range str {
		switch char {
		case '(':
			stack = append(stack, '(')
		case '[':
			stack = append(stack, '[')
		case ')':
			if (len(stack) == 0) || (stack[len(stack)-1] != '(') {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if (len(stack) == 0) || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
