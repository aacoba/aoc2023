package main

import (
	"fmt"
	"gitlab.com/auke/aoc2023/util"
	"strconv"
	"unicode"
)

func main() {
	lines, err := util.ReadFileToStringSlice("/Users/abakker/Code/aoc2023/days/03_Gear_Ratios/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Sum %d\n", ScanGearRatios(lines))

}

func ScanGearRatios(lines []string) int {
	lineMap := make(map[int][]rune, len(lines))
	for idx, line := range lines {
		lineMap[idx] = []rune(line)
	}

	sum := 0

	gears := GetGears(lines)
	for _, gear := range gears {
		numbers := make([]string, 0)
		prevLine, prevLineExists := lineMap[gear.line-1]
		if prevLineExists {
			numbers = append(numbers, ScanLineForNumbers(prevLine, gear.char)...)
		}

		numbers = append(numbers, ScanLineForNumbers(lineMap[gear.line], gear.char)...)
		nextLine, nextLineExists := lineMap[gear.line+1]
		if nextLineExists {
			numbers = append(numbers, ScanLineForNumbers(nextLine, gear.char)...)
		}

		if len(numbers) == 2 {
			val1, err := strconv.Atoi(numbers[0])
			if err != nil {
				panic(err)
			}
			val2, err := strconv.Atoi(numbers[1])
			if err != nil {
				panic(err)
			}
			sum += val1 * val2
		}
	}
	return sum
}

func ScanLineForNumbers(line []rune, position int) []string {
	scanPosition := position - 1

	numbers := make([]string, 0)
	for {
		if scanPosition-position > 1 {
			break
		}
		res := ScanLineForNumberAtPosition(line, scanPosition)
		if len(res) > 0 {
			numbers = append(numbers, string(res))
			scanPosition += 2
		} else {
			scanPosition += 1
		}
	}
	return numbers
}

func ScanLineForNumberAtPosition(line []rune, position int) []rune {
	if len(line) < 1 {
		return nil
	}

	if position < 0 || position > len(line) {
		return nil
	}
	result := make([]rune, 0)

	// Scan backwards, if result > 0, invert result
	// then scan forward at position+1

	backwardResult := ScanLine(line, position, true)
	// Nothing found, also at position
	if len(backwardResult) < 1 {
		return nil
	}
	// append to result in reverse order
	for idx := range backwardResult {
		result = append(result, backwardResult[len(backwardResult)-1-idx])
	}
	forwardResult := ScanLine(line, position+1, false)
	result = append(result, forwardResult...)

	return result
}

func ScanLine(line []rune, position int, backwards bool) []rune {
	result := make([]rune, 0)
	if position > len(line)-1 || position < 0 {
		return nil
	}
	val := line[position]
	if unicode.IsDigit(val) {
		result = append(result, val)
		nextPosition := position + scanDirection(backwards)
		x := ScanLine(line, nextPosition, backwards)
		if len(x) > 0 {
			result = append(result, x...)
		}
	}
	return result
}

func scanDirection(b bool) int {
	if b {
		return -1
	} else {
		return 1
	}
}

func GetGears(lines []string) []coordinate {
	coordinates := make([]coordinate, 0)
	for lidx, line := range lines {
		for cidx, r := range line {
			if r == '*' {
				coordinates = append(coordinates, coordinate{
					line: lidx,
					char: cidx,
				})
			}
		}
	}
	return coordinates
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

func (c coordinate) GetNeighbors() []coordinate {
	b := box{c1: c, c2: c}
	return b.GetNeighborCoordinates()
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
