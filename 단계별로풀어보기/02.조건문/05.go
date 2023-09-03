package main

import "fmt"

func main() {
	var h, m int
	fmt.Scanln(&h, &m)

	if m < 45 {
		if h == 0 {
			h = 23
		} else {
			h -= 1
		}
		m += 15
	} else {
		m -= 45
	}
	fmt.Println(h, m)
}
