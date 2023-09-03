package main

import "fmt"

func main() {
	var a, b, c, reward int
	fmt.Scan(&a, &b, &c)

	if a == b && a == c {
		reward = 10000 + a*1000
	} else if (a == b) || (a == c) {
		reward = 1000 + a*100
	} else if b == c {
		reward = 1000 + b*100
	} else {
		reward = Max(a, b, c) * 100
	}
	fmt.Println(reward)
}

func Max(a, b, c int) int {
	if a >= b {
		if c >= a {
			return c
		} else {
			return a
		}
	} else {
		if c >= b {
			return c
		} else {
			return b
		}
	}
}
