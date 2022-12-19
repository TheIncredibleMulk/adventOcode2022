package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {

	var sacks []string

	priorityTable := make(map[string]int)
	

	fmt.Println("advent of code 2022 day 3")
	fmt.Println("-------------------------------")

	readFile, err := os.Open("day3_sample_input.txt")
	// readFile, err := os.Open("day3_puzzle_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		sacks = append(sacks, fileScanner.Text())
	}
	readFile.Close()

	// Part 1
	fmt.Println(sacks)
	for _, v := range sacks {
		halfLen := len(v) / 2
		containers := SplitSubN(v, halfLen)
		for _, v := range containers[0] {
			if strings.ContainsRune(containers[1], v) {
				fmt.Printf(string(v))
			}
		}
		fmt.Println()
	}
}

func SplitSubN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}
