package main

import "fmt"

func main() {
	var total, numTotalProducts, price, numProduct int
	fmt.Scan(&total)
	fmt.Scan(&numTotalProducts)
	sum := 0

	for i := 0; i < numTotalProducts; i++ {
		fmt.Scan(&price, &numProduct)
		sum += (price * numProduct)
	}
	if total == sum {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
