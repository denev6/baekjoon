// 6864KB, 88ms

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var word string
	fmt.Fscan(reader, &word)

	word = strings.ToUpper(word)
	dict := CountChar(word)
	char := GetMostFrequent(dict)
	fmt.Fprintf(writer, "%c", char)
}

func CountChar(word string) map[rune]int {
	dict := make(map[rune]int)
	for _, char := range word {
		dict[char] += 1
	}
	return dict
}

func GetMostFrequent(dict map[rune]int) rune {
	var maxChar rune
	maxFrequency := 0
	for char, frequency := range dict {
		if maxFrequency < frequency {
			maxChar = char
			maxFrequency = frequency
		} else if maxFrequency == frequency {
			maxChar = '?'
		}
	}
	return maxChar
}
