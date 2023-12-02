package main

import (
	"fmt"
	"gitlab.com/auke/aoc2023/util"
	"strconv"
	"unicode"
)

func main() {
	lines, err := util.ReadFileToStringSlice("/Users/abakker/Code/aoc2023/days/1/input.txt")
	if err != nil {
		panic(err)
	}
	results := make([]int, len(lines))
	sum := 0
	for idx, l := range lines {
		digits, err := GetFirstAndLastDigit(l)
		if err != nil {
			panic(err)
		}
		results[idx] = digits
		sum = sum + digits
	}

	fmt.Printf("Result %d", sum)
}

func GetFirstAndLastDigit(s string) (int, error) {
	first, err := ScanFirstDigit(s)
	if err != nil {
		return 0, fmt.Errorf("cant get first digit %w", err)
	}
	last, err := ScanLastDigit(s)
	if err != nil {
		return 0, fmt.Errorf("cant get last digit %w", err)
	}
	return strconv.Atoi(fmt.Sprintf("%s%s", first, last))
}

func ScanFirstDigit(s string) (string, error) {
	for _, c := range []rune(s) {
		if unicode.IsDigit(c) {
			return string(c), nil
		}
	}
	return "", fmt.Errorf("no digit found")
}

func ScanLastDigit(s string) (string, error) {
	runes := []rune(s)
	for idx, _ := range runes {
		c := runes[len(runes)-1-idx]
		if unicode.IsDigit(c) {
			return string(c), nil
		}
	}
	return "", fmt.Errorf("no digit found")
}
