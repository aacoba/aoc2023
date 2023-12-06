package part1

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var lineRegex = regexp.MustCompile(`Card *\d+: (?P<t>[\d\s]+) \| (?P<y>[\d\s]+)`)

func CalculateScoreCards(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += GetLineScore(line)
	}
	return sum
}

func GetLineScore(line string) int {
	x := lineRegex.FindStringSubmatch(line)
	winningNumbers := SplitNumberString(x[1])
	myNumbers := SplitNumberString(x[2])
	fmt.Printf("%v, %v\n", winningNumbers, myNumbers)

	value := 0
	for _, n := range myNumbers {
		if slices.Contains(winningNumbers, n) {
			if value == 0 {
				value = 1
			} else {
				value = value * 2
			}

		}

	}
	return value
}

func SplitNumberString(s string) []int {
	split := strings.Split(s, " ")
	result := make([]int, 0)
	for _, n := range split {
		if n == "" {
			continue
		}
		i, err := strconv.Atoi(n)
		if err != nil {
			panic(fmt.Errorf("could not process number spring %s, %w", s, err))
		}
		result = append(result, i)
	}
	return result
}
