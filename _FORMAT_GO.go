package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)
var b strings.Builder

func main() {

}

func ReadLine() string {
	str, _ := reader.ReadString('\n')
	return strings.TrimRight(str, "\n")
}

func ReadValues() []string {
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	return strings.Split(str, " ")
}

/*
defer writer.Flush()
fmt.Fprint(writer, "")

fmt.Fscan(reader, &n)

sc.Scan()       
input := sc.Text()

b.WriteString()
*/