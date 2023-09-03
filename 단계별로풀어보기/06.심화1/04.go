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

	var word string
	fmt.Fscan(reader, &word)
	if isPalindrome(word) {
		fmt.Fprint(writer, 1)
	} else {
		fmt.Fprint(writer, 0)
	}

}

func isPalindrome(word string) bool {
	for i, j := 0, len(word)-1; i < len(word)/2; i, j = i+1, j-1 {
		if word[i] != word[j] {
			return false
		}
	}
	return true
}
