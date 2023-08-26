package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var str1, str2 string
	var num1, num2 int
	fmt.Fscan(reader, &str1, &str2)

	str1 = ReverseString(str1)
	str2 = ReverseString(str2)

	num1, _ = strconv.Atoi(str1)
	num2, _ = strconv.Atoi(str2)

	if num1 > num2 {
		fmt.Fprint(writer, num1)
	} else {
		fmt.Fprint(writer, num2)
	} // 같은 경우는 없음
}

func ReverseString(str string) string {
	var res string
	for _, char := range str {
		res = string(char) + res
	}
	return res
}
