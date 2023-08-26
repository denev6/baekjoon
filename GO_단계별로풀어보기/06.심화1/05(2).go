// 6832KB, 60ms

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
	alphabets := CountChar(word)
	char := GetMostFrequent(alphabets)
	fmt.Fprintf(writer, "%c", char)
}

func CountChar(word string) []int {
	alphabets := make([]int, 26)
	for _, char := range word {
		alphabets[char-'A'] += 1
	}
	return alphabets
}

func GetMostFrequent(alphabets []int) rune {
	var maxChar rune
	maxFrequency := 0
	for idx, frequency := range alphabets {
		if maxFrequency < frequency {
			maxChar = rune(idx)
			maxFrequency = frequency
		} else if maxFrequency == frequency {
			maxChar = '?'
		}
	}
	if maxChar != '?' {
		maxChar += 'A'
	}
	return maxChar
}
