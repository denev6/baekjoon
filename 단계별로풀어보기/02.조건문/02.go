package main

import "fmt"

func main() {
	var grade int
	fmt.Scan(&grade)

	if 90 <= grade && grade <= 100 {
		fmt.Println("A")
	} else if 80 <= grade && grade < 90 {
		fmt.Println("B")
	} else if 70 <= grade && grade < 80 {
		fmt.Println("C")
	} else if 60 <= grade && grade < 70 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
