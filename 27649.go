package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var b strings.Builder

func main() {
	str, _ := reader.ReadString('\n')
	str = strings.TrimRight(str, "\n")

	for i := 0; i < len(str); i++ {
		char := str[i]
		switch char {
		case '<':
			fallthrough
		case '>':
			fallthrough
		case '(':
			fallthrough
		case ')':
			b.WriteString(" ")
			b.WriteString(string(char))
			b.WriteString(" ")
		case '&':
			fallthrough
		case '|':
			// 조건: & 또는 |가 홀수 번 연속으로 주어지지 않습니다.
			if str[i+1] == char {
				b.WriteString(" ")
				b.WriteString(string(char))
				b.WriteString(string(char))
				b.WriteString(" ")
				i++
			}
		default:
			b.WriteString(string(char))
		}
	}
	words := strings.Split(b.String(), " ")

	b.Reset()
	b.WriteString(words[0])
	for _, word := range words[1:] {
		if len(word) != 0 {
			b.WriteString(" ")
			b.WriteString(word)
		}
	}
	fmt.Print(strings.TrimSpace(b.String()))
}
