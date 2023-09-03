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

	/*
		MenOfPassion(A[], n) {
			sum <- 0;
			for i <- 1 to n - 2
				for j <- i + 1 to n - 1
					for k <- j + 1 to n
						sum <- sum + A[i] × A[j] × A[k]; # 코드1
			return sum;
		}

		- i, j, k가 n에 따라 중첩으로 반복하므로 O(n^3)
		- 실행 횟수: [Latex] 
			\sum_{i=1}^{n-2}(\sum_{j=i+1}^{n-1}(\sum_{k=j+1}^{n}1))
	*/
	
	// 풀이 참고: https://rightbellboy.tistory.com/206
	cnt := (n - 2) * (n - 1) * n / 6
	fmt.Fprintf(writer, "%d\n%d", cnt, 3)
}

/*
func MenOfPassion(n int) int {
	cnt := 0
	for i := 1; i <= n-2; i++ {
		for j := i+1; j <= n-1; j++ {
			for k := j+1; k <= n; k++ {
				cnt += 1
			}
		}
	}
	return cnt 
}

*/
