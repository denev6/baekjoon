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

	var numTotal, numAward, score int
	fmt.Fscanln(reader, &numTotal, &numAward)

	scores := make([]int, numTotal)
	for i := 0; i < numTotal; i++ {
		fmt.Fscan(reader, &score)
		scores[i] = score
	}
	scores = sort(scores)
	cutLine := numTotal - numAward
	fmt.Fprint(writer, scores[cutLine])
}

func sort(arr []int) []int {
	// simple quick sort
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var leftArr, rightArr []int
	for _, n := range arr[1:] {
		if n <= pivot {
			leftArr = append(leftArr, n)
		} else {
			rightArr = append(rightArr, n)
		}
	}
	tmpArr := append(sort(leftArr), pivot)
	arr = append(tmpArr, sort(rightArr)...)
	return arr
}
