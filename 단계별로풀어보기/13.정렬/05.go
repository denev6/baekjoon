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

	var lenArr, tmp int
	fmt.Fscan(reader, &lenArr)

	var arr [10001]int

	for i := 0; i < lenArr; i++ {
		fmt.Fscan(reader, &tmp)
		arr[tmp] += 1
	}

	// 실제 정렬하지 않고 오름차순으로 출력
	for idx, n := range arr {
		if n != 0 {
			for i := 0; i < n; i++ {
				fmt.Fprintln(writer, idx)
			}
		}
	}
}
