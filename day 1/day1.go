package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var elfCalories []int

func main() {
	fmt.Println("advent of code 2022 day 1")

	readFile, err := os.Open("day1_puzzle_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var calories int
	for fileScanner.Scan() {
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
	}
	fmt.Printf("%+v", elfCalories)

	readFile.Close()

	// Find the elf carrying the most calories and print how many calories it is.
	var maxCalories [2]int
	for k, v := range elfCalories {
		if v > maxCalories[1] {
			maxCalories[0] = k
			maxCalories[1] = v
		}
	}
	fmt.Println("Total Calories Carried = ", maxCalories)

	// Find the top three Elves carrying the most Calories. How many Calories are those Elves carrying in total?
	sortedElfCalories := elfCalories
	sort.Ints(sortedElfCalories)
	top3 := sortedElfCalories[len(sortedElfCalories)-1] + sortedElfCalories[len(sortedElfCalories)-2] + sortedElfCalories[len(sortedElfCalories)-3]
	fmt.Println("Top 3 = ", top3)
}
