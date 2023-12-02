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
