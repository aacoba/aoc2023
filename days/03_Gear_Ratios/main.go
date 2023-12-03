package main

import (
	"fmt"
	"gitlab.com/auke/aoc2023/util"
	"strconv"
	"unicode"
	"unicode/utf8"
)

var dotRune, _ = utf8.DecodeRuneInString(".")

func main() {
	lines, err := util.ReadFileToStringSlice("/Users/abakker/Code/aoc2023/days/03_Gear_Ratios/input.txt")
	if err != nil {
		panic(err)
	}

	numberBoxes := make([]box, 0)
	for idx, l := range lines {
		numberBoxes = append(numberBoxes, GetNumberBoxes(l, idx)...)
	}

	boxesWithSymbolNeighbor := make([]box, 0)
	sum := 0

	for _, numberBox := range numberBoxes {
		content, err := getBoxContent(lines, numberBox)
		if err != nil {
			panic(fmt.Errorf("cant get value of box %v, %w", numberBox, err))
		}
		if hasSymbolNeighbor(lines, numberBox) {
			boxesWithSymbolNeighbor = append(boxesWithSymbolNeighbor, numberBox)
			val, err := strconv.Atoi(content)
			if err != nil {
				panic(fmt.Errorf("box content could not be converted to number %v, %w", content, err))
			}
			sum += val
			fmt.Printf("True! %v\n", content)
		} else {
			fmt.Printf("False! %v\n", content)
		}
	}

	fmt.Printf("Sum: %d", sum)
}

func hasSymbolNeighbor(lines []string, b box) bool {
	neighbors := b.GetNeighborCoordinates()
	for _, nb := range neighbors {
		if CheckIfSymbolAtCoordinate(lines, nb) {
			return true
		}
	}
	return false
}

func getBoxContent(lines []string, b box) (string, error) {
	// let's assume we only have single line boxes now
	firstChar := min(b.c1.char, b.c2.char)
	lastChar := max(b.c1.char, b.c2.char)
	output := make([]rune, 0)
	for i := firstChar; i <= lastChar; i++ {
		coord := coordinate{line: b.c1.line, char: i}
		val, err := getValueAtCoordinate(lines, coord)
		if err != nil {
			return "", fmt.Errorf("cant get value at coordinate %v, %w", coord, err)
		}
		output = append(output, val)
	}
	return string(output), nil
}
func getValueAtCoordinate(lines []string, c coordinate) (rune, error) {
	// if we ever get negatives, remove this
	if c.line < 0 {
		return '\000', fmt.Errorf("out of bounds")
	}
	if c.char < 0 {
		return '\000', fmt.Errorf("out of bounds")
	}

	if c.line >= len(lines) {
		return '\000', fmt.Errorf("out of bounds")
	}
	l := lines[c.line]
	if c.char >= len(l) {
		return '\000', fmt.Errorf("out of bounds")
	}
	return []rune(l)[c.char], nil
}

func CheckIfSymbolAtCoordinate(lines []string, c coordinate) bool {
	val, err := getValueAtCoordinate(lines, c)
	if err != nil {
		return false
	}
	return IsSymbol(val)
}

func GetNumberBoxes(line string, lineId int) []box {
	digitBoxes := make([]box, 0)
	start := -1
	for idx, char := range line {
		if unicode.IsDigit(char) {
			// not measuring a box yet
			if start < 0 {
				// start measuring a box
				start = idx
			}
		} else {
			if start >= 0 {
				digitBoxes = append(digitBoxes, box{coordinate{lineId, start}, coordinate{lineId, idx - 1}})
				start = -1
			}
		}
	}
	// close of last potential box
	if start >= 0 {
		digitBoxes = append(digitBoxes, box{coordinate{lineId, start}, coordinate{lineId, len(line) - 1}})
	}

	return digitBoxes
}

func IsSymbol(r rune) bool {
	if unicode.IsDigit(r) {
		return false
	}
	if r == '.' {
		return false
	}
	return true

}

type coordinate struct {
	line int
	char int
}

type box struct {
	c1 coordinate
	c2 coordinate
}

func (b box) GetNeighborCoordinates() []coordinate {
	linemin := min(b.c1.line, b.c2.line)
	linemax := max(b.c1.line, b.c2.line)
	charmin := min(b.c1.char, b.c2.char)
	charmax := max(b.c1.char, b.c2.char)

	coordinates := make([]coordinate, 0)

	// lines
	for l := linemin - 1; l <= linemax+1; l++ {
		for c := charmin - 1; c <= charmax+1; c++ {
			coord := coordinate{line: l, char: c}
			if !b.isCoordinateInBox(coord) { // Dont include coordinates inside the box
				coordinates = append(coordinates, coord)
			}

		}
	}
	return coordinates
}

func (b box) isCoordinateInBox(c coordinate) bool {
	linemin := min(b.c1.line, b.c2.line)
	linemax := max(b.c1.line, b.c2.line)
	charmin := min(b.c1.char, b.c2.char)
	charmax := max(b.c1.char, b.c2.char)
	if c.line >= linemin && c.line <= linemax && c.char >= charmin && c.char <= charmax {
		return true
	}
	return false
}
