package part2

import (
	"fmt"
	"gitlab.com/auke/aoc2023/days/04_Scratchcards/part1"
	"regexp"
	"slices"
	"strconv"
)

var lineRegex = regexp.MustCompile(`Card *(\d+): (?P<t>[\d\s]+) \| (?P<y>[\d\s]+)`)

type ScoreCards map[int]*ScoreCard

func (c ScoreCards) CalculateCardNumbers() int {
	count := 0
	keys := make([]int, 0)
	for k, _ := range c {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for _, k := range keys {
		card := c[k]
		for _, rc := range card.GetResultingCards() {
			c[rc].Count += card.Count
		}
		count += card.Count
	}

	return count
}

type ScoreCard struct {
	Id             int
	MyNumbers      []int
	WinningNumbers []int
	Count          int
}

func (s ScoreCard) GetMatchCount() int {
	c := 0
	for _, n := range s.MyNumbers {
		if slices.Contains(s.WinningNumbers, n) {
			c += 1
		}
	}
	return c
}

func (s ScoreCard) GetResultingCards() []int {
	winningCount := s.GetMatchCount()
	result := make([]int, 0)
	for i := 1; i <= winningCount; i++ {
		result = append(result, s.Id+i)
	}
	return result
}

func LoadScoreCards(lines []string) ScoreCards {
	cards := make(ScoreCards, len(lines))
	for _, l := range lines {
		scoreCard := ProcessScoreCardLine(l)
		cards[scoreCard.Id] = &scoreCard
	}
	return cards
}

func ProcessScoreCardLine(line string) ScoreCard {
	x := lineRegex.FindStringSubmatch(line)
	cardId, err := strconv.Atoi(x[1])
	if err != nil {
		panic(fmt.Errorf("cant parse cardId for line %s, %w", line, err))
	}
	winningNumbers := part1.SplitNumberString(x[2])
	myNumbers := part1.SplitNumberString(x[3])
	return ScoreCard{
		Id:             cardId,
		MyNumbers:      myNumbers,
		WinningNumbers: winningNumbers,
		Count:          1,
	}
}
