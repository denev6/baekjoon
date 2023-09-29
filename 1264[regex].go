package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	re := regexp.MustCompile("[aeiou]")
	for {
		sc.Scan()
		input := strings.ToLower(sc.Text())
		if input == "#" {
			return
		}
		fmt.Println(len(re.FindAllString(input, -1)))
	}
}
