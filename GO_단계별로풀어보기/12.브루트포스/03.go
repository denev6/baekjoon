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
	var a, b, c, d, e, f, x int
	fmt.Fscan(reader, &a, &b, &c, &d, &e, &f)

	/*
		   adx + bdy = cd
		-  adx + aey = af
		-------------------
		bdy - aey = cd - af

		So, y = (cd - af) / (bd - ae)

		방정식을 만족하는 (x, y)가 유일하게 존재 (x, y는 정수)
		-> bd - ae != 0

		---

		ax + by = c
		x = (c - by) / a

		OR

		dx + ey = f
		x = (f - ey) / d
	*/

	y := (c*d - a*f) / (b*d - a*e)

	if a == 0 {
		x = (f - e*y) / d
	} else {
		x = (c - b*y) / a
	}

	fmt.Fprint(writer, x, y)
}
