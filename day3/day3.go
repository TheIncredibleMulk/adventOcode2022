package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("advent of code 2022 day 3")

	readFile, err := os.Open("day3_sample_input.txt")
	// readFile, err := os.Open("day3_puzzle_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	readFile.Close()

}
