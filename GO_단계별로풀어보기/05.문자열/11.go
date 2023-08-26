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

	for {
		line, err := reader.ReadString('\n')
		fmt.Fprint(writer, line)

		if err != nil {
			return
		}
	}
}


/*

io.Copy(os.Stdout, os.Stdin)

*/
