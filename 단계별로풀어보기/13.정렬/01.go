package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var lenArr, tmp int
	fmt.Fscan(reader, &lenArr)

	arr := make([]int, lenArr)

	for i := 0; i < lenArr; i++ {
		fmt.Fscan(reader, &tmp)
		arr[i] = tmp
	}

	arr = quickSort(arr, lenArr)
	for _, n := range arr {
		fmt.Fprintf(writer, "%d\n", n)
	}
}

func quickSort(arr []int, lenArr int) []int {
	sorted := SortedArr{arr}
	sorted.Sort(0, lenArr-1)
	return sorted.GetSorted()
}

type SortedArr struct {
	arr []int
}

func (s *SortedArr) Sort(start, end int) {
	if end <= start {
		return
	}

	pivotIdx := start
	pivot := s.arr[start]
	low := start + 1
	high := end

	for low <= high {
		for low <= high && s.arr[low] <= pivot {
			low += 1
		}
		for start < high && pivot <= s.arr[high] {
			high -= 1
		}

		if high < low {
			s.arr[high], s.arr[pivotIdx] = s.arr[pivotIdx], s.arr[high]
		} else {
			s.arr[low], s.arr[high] = s.arr[high], s.arr[low]
		}
	}
	s.Sort(start, high-1)
	s.Sort(high+1, end)
}

func (s *SortedArr) GetSorted() []int {
	return s.arr
}
