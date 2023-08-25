package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var numScore int
	fmt.Fscanln(reader, &numScore)

	scores := GetFloats(numScore)
	bestScore := Max(scores)

	for idx, score := range scores {
		scores[idx] = score / bestScore * 100
	}
	mean := Mean(scores)
	fmt.Fprint(writer, mean)
}

func GetFloats(numFloats int) []float64 {
	floats := make([]float64, numFloats)
	line, _ := reader.ReadString('\n')
	strFloats := strings.Split(line, " ")

	for idx, str := range strFloats {
		floats[idx], _ = strconv.ParseFloat(strings.TrimSpace(str), 64)
	}
	return floats
}

func Max(floatSlice []float64) float64 {
	max := 0.0
	for _, float := range floatSlice {
		if float > max {
			max = float
		}
	}
	return max
}

func Mean(floatSlice []float64) float64 {
	sum := 0.0
	numFloat := len(floatSlice)
	for _, float := range floatSlice {
		sum += float
	}
	mean := sum / float64(numFloat)
	return mean
}
