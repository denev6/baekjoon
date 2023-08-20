package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	var b string
	fmt.Scan(&a)
	fmt.Scan(&b)
	
	res := 0
	value := 100
	var tmp int
	process := []int{}
	
	for _, char := range b {
		tmp, _ = strconv.Atoi(string(char))
		tmp *= a 
		process = append(process, tmp)
		
		res += tmp * value
		value /= 10
	}

	for i := 2; i >= 0; i-- {
		fmt.Println(process[i])
	}

	fmt.Println(res)
}