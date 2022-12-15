package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type winningStat struct {
	num    int
	kind   string
	points int
}

var decoder = make(map[string]string)

var rounds [][]string

var stats []winningStat

func main() {

	fmt.Println("advent of code 2022 day 2")

	// readFile, err := os.Open("day2_sample_input.txt")
	readFile, err := os.Open("day2_puzzle_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	// create map of codes to wins
	decoder["A"] = "rock"
	decoder["X"] = "rock"
	decoder["B"] = "paper"
	decoder["Y"] = "paper"
	decoder["C"] = "scissor"
	decoder["Z"] = "scissor"
	
	// point values for wins
	drawPoints := 3
	winPoints := 6
	lossPoints := 0
	rockPoints := 1
	paperPoints := 2
	scissorPoints := 3
	var points int

	for fileScanner.Scan() {
		moves := strings.Split(fileScanner.Text(), " ")
		rounds = append(rounds, moves)
	}

	readFile.Close()

	points = 0
	for i, v := range rounds {
		//  Compare each round and Figure out if we won or lost.
		opponent := decoder[v[0]]
		you := decoder[v[1]]
		var stat winningStat
		stat.num = i
		stat.kind = you

		switch you {
		case "rock":
			switch opponent {
			case "rock":
				stat.points = drawPoints + rockPoints
			case "paper":
				stat.points = lossPoints + rockPoints
			case "scissor":
				stat.points = winPoints + rockPoints

			}
		case "paper":
			switch opponent {
			case "rock":
				stat.points = winPoints + paperPoints
			case "paper":
				stat.points = drawPoints + paperPoints
			case "scissor":
				stat.points = lossPoints + paperPoints
			}
		case "scissor":
			switch opponent {
			case "rock":
				stat.points = lossPoints + rockPoints
			case "paper":
				stat.points = winPoints + scissorPoints
			case "scissor":
				stat.points = drawPoints + scissorPoints
			}

		}
		stats = append(stats, stat)
		points = points + stat.points
	}
	// for _, v := range stats {
	// 	points = v.points + points
	// }
	fmt.Printf("%+v\n", points)
}
