package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var numLog int
	var name, state string

	log := make(map[string]struct{})

	fmt.Fscan(reader, &numLog)
	for i := 0; i < numLog; i++ {
		fmt.Fscan(reader, &name, &state)

		if state == "enter" {
			log[name] = struct{}{}
			// 빈 구조체를 사용하면 메모리를 아낄 수 있음
		} else if state == "leave" {
			delete(log, name)
		}
	}

	names := make([]string, len(log))
	i := 0
	for name, _ := range log {
		names[i] = name
		i++
	}

	sort.Slice(names, func(i, j int) bool {
		return names[i] > names[j]
	})

	for _, name := range names {
		fmt.Fprintln(writer, name)
	}
}
