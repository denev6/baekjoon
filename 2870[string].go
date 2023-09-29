package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	var res []string
	re := regexp.MustCompile("[0-9]+")

	for i := 0; i < numInput; i++ {
		sc.Scan()
		input = sc.Text()
		for _, n := range re.FindAllString(input, -1) {
			tmp := strings.TrimLeft(n, "0")
			if tmp == "" {
				tmp = "0"
			}
			res = append(res, tmp)
		}
	}

	sort.Slice(res, func(i, j int) bool {
		if len(res[i]) != len(res[j]) {
			return len(res[i]) < len(res[j])
		}
		return res[i] < res[j]
	})

	b.WriteString(res[0])
	for _, n := range res[1:] {
		b.WriteString("\n")
		b.WriteString(n)
	}
	fmt.Print(b.String())
}
