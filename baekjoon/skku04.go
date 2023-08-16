package main

import (
	"bufio"
	"fmt"
	"os"
)

// 시간 초과된 풀이: O(N^2)
func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// Inputs
	var numParts int
	fmt.Fscanln(reader, &numParts)

	parts := make([]int, numParts)
	for i := 0; i < numParts; i++ {
		fmt.Fscan(reader, &parts[i])
	}

	res := 0

	for i := 1; i < numParts; i++ {
		head := Sum(parts[:i])
		if head*2 >= Sum(parts[i:]) {
			break
		}
		for j := numParts - 1; j > i; j-- {
			thorax := Sum(parts[i:j])
			abdomen := Sum(parts[j:])

			if thorax > abdomen && abdomen > head {
				res += 1
			} else if abdomen >= thorax {
				break
			}
		}
	}

	fmt.Fprintln(writer, res)
}

func Sum(numbers []int) (total int) {
	for i, _ := range numbers {
		total += numbers[i]
	}
	return total
}
