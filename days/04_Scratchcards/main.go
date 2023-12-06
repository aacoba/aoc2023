package main

import (
	"fmt"
	"gitlab.com/auke/aoc2023/days/04_Scratchcards/part2"
	"gitlab.com/auke/aoc2023/util"
)

func main() {
	lines, err := util.ReadFileToStringSlice("/Users/abakker/Code/aoc2023/days/04_Scratchcards/input.txt")
	if err != nil {
		panic(err)
	}

	//sum := part1.CalculateScoreCards(lines)
	//fmt.Printf("Sum %d", sum)

	scorecard := part2.LoadScoreCards(lines)
	fmt.Printf("%v\n", scorecard.CalculateCardNumbers())
}
