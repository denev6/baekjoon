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

	var target, res int
	fmt.Fscan(reader, &target)

	for i := GetMinRange(target); i < target; i++ {
		if GetDecomposed(i) == target {
			res = i
			break
		}
	}
	fmt.Fprint(writer, res)
}

func GetMinRange(target int) int {
	sum := 0
	for i := target; i > 0; i /= 10 {
		sum += 9
	}
	return target - sum
}

func GetDecomposed(num int) int {
	sum := num
	for i := num; i > 0; i /= 10 {
		sum += i % 10
	}
	return sum
}
