package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type round struct {
}

var decoder = make(map[string]string)

var rounds [][]string

func main() {
	fmt.Println("advent of code 2022 day 2")

	readFile, err := os.Open("day2_puzzle_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	// create map of codes to wins
	decoder["a"] = "rock"
	decoder["x"] = "rock"
	decoder["b"] = "paper"
	decoder["y"] = "paper"
	decoder["c"] = "scissor"
	decoder["z"] = "scissor"

	for fileScanner.Scan() {
		moves := strings.Split(fileScanner.Text(), " ")
		rounds = append(rounds, moves)
	}

	readFile.Close()

	for _, v := range rounds {
		//  Compare each round and Figure out if we won or lost.
		fmt.Println(v)
	}
}
