package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var elfCalories []int

func main() {
	fmt.Println("advent of code 2022 day 1")

	readFile, err := os.Open("day1_sample_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var calories int
		if fileScanner.Text() == "" {
			elfCalories = append(elfCalories, calories)
			calories = 0
		} else {
			v, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			calories = calories + v
		}
		fmt.Println(elfCalories)
	}

	readFile.Close()
}
