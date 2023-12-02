package main

import (
	"fmt"
	"gitlab.com/auke/aoc2023/util"
	"regexp"
	"strconv"
	"strings"
)

var lineRe = regexp.MustCompile(`^Game (\d+): (.*)$`)
var sampleRe = regexp.MustCompile(`(\d+) (\w+)`)

func main() {
	lines, err := util.ReadFileToStringSlice("/Users/abakker/Code/aoc2023/days/02_Cube_Conundrum/input.txt")
	if err != nil {
		panic(err)
	}
	validSum := 0
	powerSum := 0
	parsedGames := make([]GameInfo, 0, len(lines))
	for _, l := range lines {
		parsed, err := parseLine(l)
		if err != nil {
			panic(err)
		}
		parsedGames = append(parsedGames, parsed)
		if isValidGame(parsed) {
			gameId, _ := strconv.Atoi(parsed.Id)
			validSum += gameId
		} else {
			//fmt.Printf("Invalid game %v\n", parsedGames)
		}
		powerSum += powerCubes(parsed)
	}
	fmt.Printf("Sum of valid gameIds %v\n", validSum)
	fmt.Printf("Powersum %v\n", powerSum)
}

func isValidGame(g GameInfo) bool {
	for _, s := range g.Samples {
		if s.RedCount > 12 {
			return false
		}
		if s.GreenCount > 13 {
			return false
		}
		if s.BlueCount > 14 {
			return false
		}
	}
	return true
}

func powerCubes(g GameInfo) int {
	minRed := 0
	minGreen := 0
	minBlue := 0
	for _, s := range g.Samples {
		if s.RedCount > minRed {
			minRed = s.RedCount
		}
		if s.GreenCount > minGreen {
			minGreen = s.GreenCount
		}
		if s.BlueCount > minBlue {
			minBlue = s.BlueCount
		}
	}
	return minRed * minGreen * minBlue
}

type GameInfo struct {
	Id      string
	Samples []Sample
}

type Sample struct {
	RedCount   int
	GreenCount int
	BlueCount  int
}

func parseLine(l string) (GameInfo, error) {
	matches := lineRe.FindAllStringSubmatch(l, -1)
	if len(matches) != 1 {
		return GameInfo{}, fmt.Errorf("expected one match but found %d. [%s]", len(matches), l)
	}
	parsedSamples, err := parseSamples(matches[0][2])
	if err != nil {
		return GameInfo{}, fmt.Errorf("could not parse sampleline %s, %w", l, err)
	}

	return GameInfo{
		matches[0][1],
		parsedSamples,
	}, nil
}

func parseSamples(l string) ([]Sample, error) {
	rawSamples := strings.Split(l, ";")
	results := make([]Sample, 0)
	for _, s := range rawSamples {
		parsed, err := parseSample(s)
		if err != nil {
			return results, fmt.Errorf("could not parse sample [%v], %w", s, err)
		}
		results = append(results, parsed)
	}
	return results, nil
}

func parseSample(l string) (Sample, error) {
	sample := Sample{}
	matches := sampleRe.FindAllStringSubmatch(l, -1)
	for _, match := range matches {
		if len(match) != 3 {
			return sample, fmt.Errorf("invalid syntax for sample %s", l)
		}
		sampleValue, err := strconv.Atoi(match[1])
		if err != nil {
			return sample, fmt.Errorf("invalid sample value %s in sample %s", match[1], l)
		}
		switch match[2] {
		case "red":
			sample.RedCount = sampleValue
		case "green":
			sample.GreenCount = sampleValue
		case "blue":
			sample.BlueCount = sampleValue
		default:
			return sample, fmt.Errorf("unknown sample type %s in sample %s", match[2], l)
		}
	}
	return sample, nil
}
