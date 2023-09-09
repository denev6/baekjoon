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

	var numNeverHeard, numNeverSeen int
	var name string

	peopleNeverHeard := make(map[string]struct{})
	var peopleDontKnow []string

	fmt.Fscan(reader, &numNeverHeard, &numNeverSeen)
	for i := 0; i < numNeverHeard; i++ {
		fmt.Fscan(reader, &name)
		peopleNeverHeard[name] = struct{}{}
	}

	for i := 0; i < numNeverSeen; i++ {
		fmt.Fscan(reader, &name)
		if _, exists := peopleNeverHeard[name]; exists {
			peopleDontKnow = append(peopleDontKnow, name)
		}
	}

	sort.Strings(peopleDontKnow)
	fmt.Fprintln(writer, len(peopleDontKnow))
	for _, name := range peopleDontKnow {
		fmt.Fprintln(writer, name)
	}
}
