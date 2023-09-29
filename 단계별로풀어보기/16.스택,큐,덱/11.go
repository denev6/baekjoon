package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var reader = bufio.NewReader(os.Stdin)
var b strings.Builder

func main() {
	var lenItems, input, flag, lenNewItems int 
	fmt.Fscan(reader, &lenItems)
	
	flags := make([]int, lenItems)
	for i := 0; i < lenItems; i++ {
		fmt.Fscan(reader, &flag)
		flags[i] = flag
	}

	qs := list.New()
	for i := 0; i < lenItems; i++ {
		fmt.Fscan(reader, &input)
		if flags[i] == 0 {
			qs.PushBack(input)
		}
	}

	fmt.Fscan(reader, &lenNewItems)
	res := make([]int, lenNewItems)

	for i := 0; i < lenNewItems; i++ {
		fmt.Fscan(reader, &input)
		qs.PushFront(input)
		res[i] = qs.Back().Value.(int)
		qs.Remove(qs.Back())
	}
	
	b.WriteString(strconv.Itoa(res[0]))
	for _, n := range res[1:] {
		b.WriteString(" ")
		b.WriteString(strconv.Itoa(n))
	}
	fmt.Print(b.String())
}
