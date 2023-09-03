package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var lenArr, tmp int
	fmt.Fscan(reader, &lenArr)

	pointerArr := make([]int, lenArr) // 정렬 순서로 변경
	valueArr := make([]int, lenArr)   // 입력 순서 유지

	for i := 0; i < lenArr; i++ {
		fmt.Fscan(reader, &tmp)
		pointerArr[i] = i
		valueArr[i] = tmp
	}

	sort.Slice(pointerArr, func(i, j int) bool {
		return valueArr[pointerArr[i]] < valueArr[pointerArr[j]]
	})

	cnt := 0
	previous := valueArr[pointerArr[0]]
	for i := 0; i < lenArr; i++ {
		if valueArr[pointerArr[i]] != previous {
			previous = valueArr[pointerArr[i]]
			cnt += 1
		}
		valueArr[pointerArr[i]] = cnt
	}

	for _, n := range valueArr {
		fmt.Fprintf(writer, "%d ", n)
	}
}
