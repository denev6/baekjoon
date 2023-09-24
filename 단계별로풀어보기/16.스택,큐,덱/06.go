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

	queue := []int{}
	var numCommand int
	fmt.Fscan(reader, &numCommand)

	for i := 0; i < numCommand; i++ {
		var command string 
		var val int
		fmt.Fscan(reader, &command)

		switch command {
		case "push":
			fmt.Fscan(reader, &val)
			queue = append(queue, val)
		case "pop":
			len_queue := len(queue)
			if len_queue > 0 {
				fmt.Fprintln(writer, queue[0])
				queue = queue[1:]
			} else {
				fmt.Fprintln(writer, -1)
			}
		case "size":
			fmt.Fprintln(writer, len(queue))
		case "empty":
			if len(queue) > 0 {
				fmt.Fprintln(writer, 0)
			} else {
				fmt.Fprintln(writer, 1)
			}
		case "front":
			len_queue := len(queue)
			if len_queue > 0 {
				fmt.Fprintln(writer, queue[0])
			} else {
				fmt.Fprintln(writer, -1)
			}
		case "back":
			len_queue := len(queue)
			if len_queue > 0 {
				fmt.Fprintln(writer, queue[len_queue-1])
			} else {
				fmt.Fprintln(writer, -1)
			}
		}
	}
}
