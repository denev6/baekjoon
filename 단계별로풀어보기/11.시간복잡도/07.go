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
	var w, b, c, n0, res int 
	fmt.Fscanln(reader, &w, &b)
	fmt.Fscanln(reader, &c)
	fmt.Fscanln(reader, &n0)

	/*
		f(n) = wn + b
		g(n) = n
		O(g(n)); 모든 n ≥ n0에 대하여 f(n) ≤ c × g(n)
	
		1) w > 0
			a) c < w
				n이 충분히 커지면(f(n)=g(n)을 지나면) 만족하지 않음
			b) w ≤ c
				n0가 만족하면 모든 n에 대하여 만족
		
		2) w = 0
			n0가 만족하면 모든 n에 대하여 만족

		3) w < 0
			n0가 만족하면 모든 n에 대하여 만족
	*/
	
	if w <= 0 || w <= c {
		if w * n0 + b <= c * n0 {
			res = 1
		} 
	}
	
	fmt.Fprint(writer, res)
}
