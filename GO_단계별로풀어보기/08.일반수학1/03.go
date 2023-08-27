package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var numTestCase, change, numCoin int
	coins := [...]int{25, 10, 5, 1}
	changeCoins := make([]int, 4)

	fmt.Fscan(reader, &numTestCase)
	for i := 0; i < numTestCase; i++ {
		fmt.Fscan(reader, &change)
		for idx, coin := range coins {
			numCoin = change / coin
			changeCoins[idx] = numCoin
			change -= numCoin * coin
		}
		PrintCoins(changeCoins)
	}
}

func PrintCoins(changeCoins []int) {
	var res string
	for _, n := range changeCoins {
		res += strconv.Itoa(n) + " "
	}
	fmt.Fprintln(writer, res)
}
