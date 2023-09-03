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

	var subject, grade string
	var score, myScore, sumScore float64
	gradeDict := map[string]float64{
		"A+": 4.5,
		"A0": 4.0,
		"B+": 3.5,
		"B0": 3.0,
		"C+": 2.5,
		"C0": 2.0,
		"D+": 1.5,
		"D0": 1.0,
		"F":  0.0,
	}

	for {
		_, err := fmt.Fscan(reader, &subject, &score, &grade)
		if err != nil {
			break
		}
		if grade == "P" {
			continue
		}
		myScore += score * gradeDict[grade]
		sumScore += score
	}
	myScore /= sumScore
	fmt.Fprintf(writer, "%f", myScore)
}
