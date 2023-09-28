package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var b strings.Builder

func main() {
	sc.Scan()
	input := sc.Text()
	numInput, _ := strconv.Atoi(input)

	words := make([]string, numInput)
	for i := 0; i < numInput; i++ {
		sc.Scan()
		words[i] = sc.Text()
	}

	sort.Slice(words, func(i, j int) bool {
		len_i, len_j := len(words[i]), len(words[j])
		if len_i == len_j {
			return words[i] < words[j]
		}
		return len_i < len_j
	})

	prev := words[0]
	b.WriteString(prev)
	for _, word := range words[1:] {
		if word != prev {
			b.WriteString("\n")
			b.WriteString(word)
			prev = word
		}
	}
	fmt.Print(b.String())
}
