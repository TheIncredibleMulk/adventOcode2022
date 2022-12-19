package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type winningStat struct {
	num         int
	kind        string
	outcome     int
	playerValue int
}

var decoder = make(map[string]string)

var rounds [][]string

var roundOneStats []winningStat

func main() {

	fmt.Println("advent of code 2022 day 2")

	// readFile, err := os.Open("day2_sample_input.txt")
	readFile, err := os.Open("day2_puzzle_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		moves := strings.Split(fileScanner.Text(), " ")
		rounds = append(rounds, moves)
	}

	readFile.Close()

	// point values for wins
	drawPoints := 3
	winPoints := 6
	lossPoints := 0
	rockPoints := 1
	paperPoints := 2
	scissorsPoints := 3

	// create map of codes to wins
	decoder["A"] = "rock"
	decoder["X"] = "rock"
	decoder["B"] = "paper"
	decoder["Y"] = "paper"
	decoder["C"] = "scissors"
	decoder["Z"] = "scissors"

	for i, v := range rounds {
		//  Compare each round and Figure out if we won or lost.
		opponentChoice := decoder[v[0]]
		playerChoice := decoder[v[1]]
		var stat winningStat
		stat.num = i
		stat.kind = playerChoice

		switch playerChoice {
		case "rock":
			switch opponentChoice {
			case "rock":
				stat.outcome = drawPoints
				stat.playerValue = rockPoints
			case "paper":
				stat.outcome = lossPoints
				stat.playerValue = rockPoints
			case "scissors":
				stat.outcome = winPoints
				stat.playerValue = rockPoints

			}
		case "paper":
			switch opponentChoice {
			case "rock":
				stat.outcome = winPoints
				stat.playerValue = paperPoints
			case "paper":
				stat.outcome = drawPoints
				stat.playerValue = paperPoints
			case "scissors":
				stat.outcome = lossPoints
				stat.playerValue = paperPoints
			}
		case "scissors":
			switch opponentChoice {
			case "rock":
				stat.outcome = lossPoints
				stat.playerValue = scissorsPoints
			case "paper":
				stat.outcome = winPoints
				stat.playerValue = scissorsPoints
			case "scissors":
				stat.outcome = drawPoints
				stat.playerValue = scissorsPoints
			}

		}
		roundOneStats = append(roundOneStats, stat)
	}
	totalOutcome := 0
	totalPlayerValue := 0
	for _, v := range roundOneStats {
		totalOutcome += v.outcome
		totalPlayerValue += v.playerValue
	}
	fmt.Println(totalOutcome, totalPlayerValue, totalOutcome+totalPlayerValue)
	fmt.Println("------------------------------------------------------------")

	// round two
	decoder["X"] = "loose"
	decoder["Y"] = "draw"
	decoder["Z"] = "win"

	var roundTwoStats []winningStat

	for i, v := range rounds {
		opponentShape := decoder[v[0]]
		playersPlay := decoder[v[1]]
		var stat winningStat
		stat.num = i
		stat.kind = playersPlay

		switch playersPlay {
		case "win":
			stat.outcome = winPoints
			switch opponentShape {
			case "rock":
				stat.playerValue = paperPoints
			case "paper":
				stat.playerValue = scissorsPoints
			case "scissors":
				stat.playerValue = rockPoints
			}

		case "draw":
			stat.outcome = drawPoints
			switch opponentShape {
			case "rock":
				stat.playerValue = rockPoints
			case "paper":
				stat.playerValue = paperPoints
			case "scissors":
				stat.playerValue = scissorsPoints

			}

		case "loose":
			stat.outcome = lossPoints
			switch opponentShape {
			case "rock":
				stat.playerValue = scissorsPoints
			case "paper":
				stat.playerValue = rockPoints
			case "scissors":
				stat.playerValue = paperPoints
			}

		}
		roundTwoStats = append(roundTwoStats, stat)

	}
	roundTwoTotalOutcome := 0
	roundTwoTotalPlayerValue := 0
	for _, v := range roundTwoStats {
		roundTwoTotalOutcome += v.outcome
		roundTwoTotalPlayerValue += v.playerValue
	}
	fmt.Println(roundTwoTotalOutcome, roundTwoTotalPlayerValue, roundTwoTotalOutcome+roundTwoTotalPlayerValue)

}
