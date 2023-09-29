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

func ReadLine() {
	sc.Scan()
	return sc.Text()
}

/*
defer writer.Flush()
fmt.Fprint(writer, "")

fmt.Fscan(reader, &n)

sc.Scan()       
input := sc.Text()

b.WriteString()
*/