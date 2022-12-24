package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type inCommonStruct struct {
	priority int
	letter   string
}

func main() {

	var sacks []string

	priorityTable := make(map[string]int)

	var inCommon []inCommonStruct
	c := 1
	for r := 'a'; r <= 'z'; r++ {
		R := unicode.ToUpper(r)
		priorityTable[string(R)] = c + 26
		priorityTable[string(r)] = c
		c++
	}
	fmt.Println("priotity Table: ", priorityTable)
	fmt.Println("-------------------------------")

	fmt.Println("advent of code 2022 day 3")
	fmt.Println("-------------------------------")

	// readFile, err := os.Open("day3_sample_input.txt")
	readFile, err := os.Open("day3_puzzle_input.txt")
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
	for _, val := range sacks {
		halfLen := len(val) / 2
		containers := SplitSubN(val, halfLen)
		commonLetter := ' '
		for _, v := range containers[0] {
			a := inCommonStruct{
				priority: priorityTable[string(v)],
				letter:   string(v),
			}
			// fmt.Println(i, string(v), string(commonLetter))
			if strings.ContainsRune(containers[1], v) {
				if commonLetter != v {
					commonLetter = v
					inCommon = append(inCommon, a)
				}
			}
		}
	}
	fmt.Println("-------------------------------")
	fmt.Printf("items in both compartments: %v\n", inCommon)

	var addedUp int
	for _, v := range inCommon {
		addedUp = addedUp + v.priority
	}
	fmt.Println("-------------------------------")
	fmt.Println(addedUp)
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
