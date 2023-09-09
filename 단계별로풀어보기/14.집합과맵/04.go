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

	var numPokemon, numQuestion int
	var pokemonName, question string

	pokemonsByName := make(map[string]int)
	pokemonsByIndex := make(map[int]string)

	fmt.Fscan(reader, &numPokemon, &numQuestion)
	for i := 1; i <= numPokemon; i++ {
		fmt.Fscan(reader, &pokemonName)
		pokemonsByName[pokemonName] = i
		pokemonsByIndex[i] = pokemonName
	}

	for i := 0; i < numQuestion; i++ {
		fmt.Fscan(reader, &question)

		if pokemonIdx, err := strconv.Atoi(question); err == nil {
			fmt.Fprintln(writer, pokemonsByIndex[pokemonIdx])
		} else {
			fmt.Fprintln(writer, pokemonsByName[question])
		}
	}
}
