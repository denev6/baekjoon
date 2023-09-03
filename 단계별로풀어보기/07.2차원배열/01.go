package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var numColumn, numRow int
	fmt.Fscan(reader, &numColumn, &numRow)

	matrixA := GetMatrix(numColumn, numRow)
	matrixB := GetMatrix(numColumn, numRow)
	matrixRes := AddMatrix(matrixA, matrixB, numColumn, numRow)
	fmt.Fprint(writer, matrixRes)
}

func GetMatrix(numColumn, numRow int) [][]int {
	var num int
	matrix := make([][]int, numColumn)
	for i := 0; i < numColumn; i++ {
		matrix[i] = make([]int, numRow)
		for j := 0; j < numRow; j++ {
			fmt.Fscan(reader, &num)
			matrix[i][j] = num
		}
	}
	return matrix
}

func AddMatrix(matrixA, matrixB [][]int, numColumn, numRow int) string {
	var matrix string
	for i := 0; i < numColumn; i++ {
		for j := 0; j < numRow; j++ {
			matrix += strconv.Itoa(matrixA[i][j]+matrixB[i][j]) + " "
		}
		if i != numColumn-1 {
			matrix += "\n"
		}
	}
	return matrix
}
