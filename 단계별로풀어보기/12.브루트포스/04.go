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
	var m, n, tmp int
	var input string  
	fmt.Fscanln(reader, &m, &n)
	
	board := make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Fscanln(reader, &input)
		board[i] = input
	}

	res := 65

	for i := 0; i <= m-8; i++ {
		for j := 0; j <= n-8; j++ {
			tmp = GetNumWrong(board, i, j)
			if tmp < res {
				res = tmp 
			}
		}
	} 
	fmt.Fprint(writer, res)
}

func GetNumWrong(board []string, indexs ...int) int {
	startCol, startRow := indexs[0], indexs[1]
	cnt := 0
	firstColor := 'B'
	secondColor := 'W' 
	var color rune 
	
	for i := startCol; i < startCol+8; i++ {
		for j := startRow; j < startRow+8; j++  {
			
			color = rune(board[i][j])
			if j % 2 == 0 {
				if color != firstColor {
					cnt += 1
				}
			} else {
				if color != secondColor {
					cnt += 1
				}
			}
		}
		firstColor, secondColor = secondColor, firstColor 
	}
	if cnt < 33 {
		// 33 = 64 / 2 + 1
		return cnt
	} 
	return 64 - cnt
}