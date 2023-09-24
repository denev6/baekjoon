// 4ms, 916KB

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

	var numPs int
	var ps string
	fmt.Fscanln(reader, &numPs)

	for i := 0; i < numPs; i++ {
		fmt.Fscanln(reader, &ps)
		
		if IsVPS(ps) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func IsVPS(ps string) bool {
	cnt := 0
	for _, char := range ps {
		if char == '(' {
			cnt++
		} else {
			cnt--
			if cnt == -1 {
				break
			}
		}
	}
	if cnt == 0 {
		return true
	} else {
		return false
	}
}