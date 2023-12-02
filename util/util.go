package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileToStringSlice(path string) ([]string, error) {
	output := make([]string, 0)
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file %w", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file, %w", err)
	}

	return output, nil

}

func ReverseString(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
