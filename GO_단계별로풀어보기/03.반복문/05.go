package main

import "fmt"

func main() {
	var numBytes int
	fmt.Scan(&numBytes)
	res := ""

	for i := 0; i < (numBytes / 4); i++ {
		res += "long "
	}
	res += "int"
	fmt.Print(res)
}
