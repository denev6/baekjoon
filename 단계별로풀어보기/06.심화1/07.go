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

	var numWord, numGroupWord int
	var word string
	fmt.Fscan(reader, &numWord)

	for i := 0; i < numWord; i++ {
		fmt.Fscan(reader, &word)
		if isGroupWord(word) {
			numGroupWord += 1
		}
	}
	fmt.Fprint(writer, numGroupWord)
}

func isGroupWord(word string) bool {
	isUsed := make([]bool, 26)
	var currnetChar rune

	for _, char := range word {
		if currnetChar != char {
			currnetChar = char
			if isUsed[char-'a'] {
				return false
			} else {
				isUsed[char-'a'] = true
			}
		}
	}
	return true
}
