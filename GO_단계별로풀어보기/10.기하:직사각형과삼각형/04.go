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

	var n int
	fmt.Fscan(reader, &n)

	// 가로: 각 층의 윗부분은 1, 가장 아래 바닥은 n
	// 따라서, 바닥의 수평한 변은 2n (1 * n + n)
	// 세로: 좌우측 변 모두 n, 따라서 2n (n + n)
	// 전체: 4n (2n + 2n)
	
	fmt.Fprint(writer, 4 * n)

}
