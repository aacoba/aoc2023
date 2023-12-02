package main

import (
	"bufio"
	"fmt"
	"gitlab.com/auke/aoc2023/util"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var digitTokens = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	lines, err := util.ReadFileToStringSlice("/Users/abakker/Code/aoc2023/days/01_Trebuchet/input.txt")
	if err != nil {
		panic(err)
	}
	results := make([]int, len(lines))
	sum := 0
	for idx, l := range lines {
		digits, err := FancyScannerGetFirstAndLastDigits(l)
		if err != nil {
			panic(err)
		}
		results[idx] = digits
		sum = sum + digits
	}

	fmt.Printf("Result %d", sum)
}

// Inspired by / stolen from by https://cs.opensource.google/go/go/+/refs/tags/go1.21.4:src/bufio/scan.go;l=395
func NumberScanner(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(r) {
			break
		}
	}

	// Scan until a digit is found or a complete word-digit is found
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		// Check if normal digit
		if unicode.IsDigit(r) {
			return i + width, data[i : i+width], nil
		}
		wordDigit := getWordDigit(data[:i+width])
		if wordDigit != "" {
			return 1, []byte(wordDigit), nil
		}
	}
	// Request more data.
	return start, nil, nil

}

func getWordDigit(b []byte) string {
	s := string(b)
	for token, value := range digitTokens {
		if strings.HasSuffix(s, token) {
			return value
		}
	}
	return ""

}

func FancyScannerGetFirstAndLastDigits(s string) (int, error) {
	digits, err := ScanStringForDigitTokens(s)
	if err != nil {
		return 0, err
	}
	if len(digits) < 1 {
		return 0, fmt.Errorf("no digits found in string %v", s)
	}
	return strconv.Atoi(fmt.Sprintf("%s%s", digits[0], digits[len(digits)-1]))

}

func ScanStringForDigitTokens(s string) ([]string, error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	output := make([]string, 0)
	scanner.Split(NumberScanner)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file, %w", err)
	}
	return output, nil
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
