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

	var roomNum int
	fmt.Fscan(reader, &roomNum)

	currentRoom := 1

	for i := 0; ; i++ {
		currentRoom += 6 * i
		if currentRoom+1 > roomNum {
			fmt.Fprint(writer, i+1)
			break
		}
	}
}
