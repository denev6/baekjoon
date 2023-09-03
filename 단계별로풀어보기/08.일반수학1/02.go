package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var num, base, tmp int
	var res string
	fmt.Fscan(reader, &num, &base)

	currentBase := 1
	for currentBase < num {
		currentBase *= base
	}
	if currentBase > num && currentBase >= base {
		currentBase /= base
	}

	for ; 0 < currentBase; currentBase /= base {
		tmp = num / currentBase
		res += string(DecimalToString(tmp))
		num -= tmp * currentBase
	}
	fmt.Fprint(writer, res)
}

func DecimalToString(dec int) rune {
	if 0 <= dec && dec <= 9 {
		return rune(dec) + '0'
	} else {
		return rune(dec-10) + 'A'
	}
}
