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

	var numWords int
	var word string // 영어 소문자와 대문자로만 이루어진 단어
	fmt.Fscan(reader, &numWords)

	for i := 0; i < numWords; i++ {
		fmt.Fscan(reader, &word)
		fmt.Fprintf(writer, "%c%c\n", word[0], word[len(word)-1])
	}
}
