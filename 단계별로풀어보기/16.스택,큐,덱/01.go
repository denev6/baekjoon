// 시간 단축을 위한 풀이
// 통과

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

	stack := []int{}
	var numCommand int
	fmt.Fscan(reader, &numCommand)

	for i := 0; i < numCommand; i++ {
		var command int
		fmt.Fscan(reader, &command)

		switch command {
		case 1:
			var val int
			fmt.Fscan(reader, &val)
			stack = append(stack, val)
		case 2:
			len_stack := len(stack)
			if len_stack > 0 {
				fmt.Fprintln(writer, stack[len_stack-1])
				stack = stack[:len_stack-1]
			} else {
				fmt.Fprintln(writer, -1)
			}
		case 3:
			fmt.Fprintln(writer, len(stack))
		case 4:
			if len(stack) > 0 {
				fmt.Fprintln(writer, 0)
			} else {
				fmt.Fprintln(writer, 1)
			}
		case 5:
			len_stack := len(stack)
			if len_stack > 0 {
				fmt.Fprintln(writer, stack[len_stack-1])
			} else {
				fmt.Fprintln(writer, -1)
			}
		}
	}
}
