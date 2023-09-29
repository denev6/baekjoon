package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var b strings.Builder

func main() {
	reader.ReadString('\n') // Skip Input
	flags := ReadValues()

	qs := list.New()
	items := ReadValues()

	for i, flag := range flags {
		if flag == "0" {
			qs.PushBack(items[i])
		}
	}

	reader.ReadString('\n') // Skip Input 
	newItems := ReadValues()
	// Reuse newItems to store results

	for i, new := range newItems {
		qs.PushFront(new)
		newItems[i] = qs.Back().Value.(string)
		qs.Remove(qs.Back())
	}

	res := newItems // 'Shallow Copy' to rename slice
	b.WriteString(res[0])
	for _, n := range res[1:] {
		b.WriteString(" ")
		b.WriteString(n)
	}
	fmt.Print(b.String())
}

func ReadValues() []string {
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	return strings.Split(str, " ")
}
