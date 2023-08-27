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

	var num int
	fmt.Fscan(reader, &num)

	if num == 1 {
		return
	}
	for i := 2; i*i <= num; i++ {
		for num%i == 0 {
			num /= i
			fmt.Fprintln(writer, i)
		}
	}
	if num != 1 {
		fmt.Fprintln(writer, num)
	}
}

/*

[실행 속도]
1.   4ms: (현재) 입력의 제곱근까지 순회하며 확인
2. 108ms: 입력값까지 순회하며 확인
3. 228ms: 모든 소수를 구한 후 순회하며 확인 

*/