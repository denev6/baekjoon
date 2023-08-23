package main

import "fmt"

func main() {
	var n, tmp, target int
	fmt.Scan(&n)
	numbers := make([]int, n)
	
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		numbers[i] = tmp
	}

	fmt.Scan(&target)
	res := 0
	
	for _, num := range numbers {
		if num == target {
			res += 1
		}
	}
	fmt.Println(res)
}
