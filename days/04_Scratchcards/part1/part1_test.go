package part1

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/auke/aoc2023/util"
	"testing"
)

func TestCalculateScoreCards(t *testing.T) {
	type test struct {
		inputFile string
		want      int
	}

	tests := []test{
		{
			"testdata/sample.txt",
			13,
		},
	}

	for _, tc := range tests {
		t.Run(tc.inputFile, func(t *testing.T) {
			lines, err := util.ReadFileToStringSlice(tc.inputFile)
			if err != nil {
				t.Fatalf("Could not read inputfile, %v", err)
			}
			assert.Equal(t, tc.want, CalculateScoreCards(lines))

		})
	}
}
