package part2

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
			30,
		},
	}

	for _, tc := range tests {
		t.Run(tc.inputFile, func(t *testing.T) {
			lines, err := util.ReadFileToStringSlice(tc.inputFile)
			if err != nil {
				t.Fatalf("Could not read inputfile, %v", err)
			}
			sc := LoadScoreCards(lines)
			got := sc.CalculateCardNumbers()
			t.Logf("Want %v got %v", tc.want, got)
			assert.Equal(t, tc.want, got)

		})
	}
}
