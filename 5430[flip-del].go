// 시간 초과

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var b strings.Builder

func main() {
	flag := true
	numCase, _ := strconv.Atoi(ReadLine())
	for i := 0; i < numCase; i++ {
		commands := ReadLine()
		numValues, _ := strconv.Atoi(ReadLine())
		strSlice := NewData(ReadLine(), numValues)

		for _, cmd := range commands {
			if cmd == 'R' {
				strSlice.Flip()
			} else if cmd == 'D' {
				flag = strSlice.DeleteFront()
				if !flag {
					break
				}
			}
		}

		if flag {
			strSlice.BuildString()
		} else {
			b.WriteString("error")
			flag = true
		}
		b.WriteString("\n")
	}
	fmt.Print(b.String())
}

func ReadLine() string {
	str, _ := reader.ReadString('\n')
	return strings.TrimRight(str, "\n")
}

type Data struct {
	values []string
	len    int
}

func NewData(str string, lenSlice int) *Data {
	if lenSlice == 0 {
		return &Data{make([]string, 0), 0}
	}
	str = strings.Trim(str, "[]")
	strSlice := strings.Split(str, ",")
	return &Data{strSlice, lenSlice}
}

func (d *Data) Flip() {
	if d.len > 2 {
		for pFront, pBack := 0, d.len-1; pFront < pBack; pFront, pBack = pFront+1, pBack-1 {
			d.values[pFront], d.values[pBack] = d.values[pBack], d.values[pFront]
		}
	}
}

func (d *Data) DeleteFront() bool {
	if d.len == 0 {
		return false
	}
	d.values = d.values[1:]
	d.len -= 1
	return true
}

func (d *Data) BuildString() {
	b.WriteString("[")
	b.WriteString(strings.Join(d.values, ","))
	b.WriteString("]")
}
