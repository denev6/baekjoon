package main

import "fmt"

func main() {
	var h, m, time int
	fmt.Scan(&h, &m)
	fmt.Scan(&time)
	time_h := time / 60
	time_m := time % 60
	h += time_h
	m += time_m

	if 60 <= m {
		h += 1
		m -= 60
	}
	if 24 <= h {
		h -= 24
	}
	fmt.Println(h, m)
}
