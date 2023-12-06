package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gitlab.com/auke/aoc2023/util"
	"testing"
)

func TestGetNumberBoxes(t *testing.T) {
	type test struct {
		name   string
		input  string
		lineid int
		want   []box
	}

	tests := []test{
		{
			name:   "test",
			input:  "...345...4.%.6*..",
			lineid: 1,
			want: []box{
				{c1: coordinate{1, 3}, c2: coordinate{line: 1, char: 5}},
				{c1: coordinate{1, 9}, c2: coordinate{line: 1, char: 9}},
				{c1: coordinate{1, 13}, c2: coordinate{line: 1, char: 13}},
			},
		},
		{
			name:   "test",
			input:  "123...7.9",
			lineid: 1,
			want: []box{
				{c1: coordinate{1, 0}, c2: coordinate{line: 1, char: 2}},
				{c1: coordinate{1, 6}, c2: coordinate{line: 1, char: 6}},
				{c1: coordinate{1, 8}, c2: coordinate{line: 1, char: 8}},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := GetNumberBoxes(tc.input, tc.lineid)
			assert.Equal(t, tc.want, result)
		})
	}
}

func TestBox_GetNeighborCoordinates(t *testing.T) {
	type test struct {
		name  string
		input box
		want  []coordinate
	}

	tests := []test{
		{
			name:  "1x1 box",
			input: box{c1: coordinate{line: 0, char: 0}, c2: coordinate{line: 0, char: 0}},
			want: []coordinate{
				{line: -1, char: -1},
				{line: -1, char: 0},
				{line: -1, char: 1},
				{line: 0, char: -1},
				{line: 0, char: 1},
				{line: 1, char: -1},
				{line: 1, char: 0}, // why is this one missing?
				{line: 1, char: 1},
			},
		},
		{
			name:  "2x2 box",
			input: box{c1: coordinate{line: 0, char: 0}, c2: coordinate{line: 1, char: 1}},
			want: []coordinate{
				{line: -1, char: -1},
				{line: -1, char: 0},
				{line: -1, char: 1},
				{line: -1, char: 2},
				{line: 0, char: -1},
				{line: 0, char: 2},
				{line: 1, char: -1},
				{line: 1, char: 2},
				{line: 2, char: -1},
				{line: 2, char: 0},
				{line: 2, char: 1},
				{line: 2, char: 2},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.input.GetNeighborCoordinates()
			assert.Equal(t, tc.want, result)
		})
	}
}

func TestIsSymbol(t *testing.T) {
	type test struct {
		input rune
		want  bool
	}

	tests := []test{{
		'*', true,
	}, {
		'.', false,
	}, {'8', false}}

	for _, tc := range tests {
		t.Run(string(tc.input), func(t *testing.T) {
			assert.Equal(t, tc.want, IsSymbol(tc.input))
		})
	}
}

func TestScanLine(t *testing.T) {
	type test struct {
		inputLine string
		position  int
		backwards bool
		want      string
	}

	tests := []test{
		{
			"..123...",
			4,
			true,
			"321",
		},

		{
			"..123...",
			2,
			false,
			"123",
		},
		{
			"..123...",
			4,
			false,
			"3",
		},
		{
			"..123...",
			5,
			false,
			"",
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Backward %v Position %d Line %s", tc.backwards, tc.position, tc.inputLine), func(t *testing.T) {
			res := ScanLine([]rune(tc.inputLine), tc.position, tc.backwards)
			assert.Equal(t, tc.want, string(res))
		})
	}
}

func TestScanLineForNumberAtPosition(t *testing.T) {
	type test struct {
		inputLine string
		position  int
		want      string
	}

	tests := []test{
		{
			"...1234...",
			3,
			"1234",
		},
		{
			"...1234...",
			4,
			"1234",
		},
		{
			"...1234...",
			6,
			"1234",
		},
		{
			"...1234...",
			7,
			"",
		},
		{
			"...1234...",
			2,
			"",
		},

		{
			"...1234...",
			12,
			"",
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Input %s, Position %d", tc.inputLine, tc.position), func(t *testing.T) {
			res := ScanLineForNumberAtPosition([]rune(tc.inputLine), tc.position)
			assert.Equal(t, tc.want, string(res))
		})
	}
}

func TestScanLineForNumbers(t *testing.T) {
	type test struct {
		inputLine string
		position  int
		want      []string
	}

	tests := []test{
		{
			"...123.456...",
			3,
			[]string{"123"},
		},
		{
			"...123.456...",
			5,
			[]string{"123"},
		},
		{
			"...123.456...",
			6,
			[]string{"123", "456"},
		},
		{
			"...123.456...",
			7,
			[]string{"456"},
		},
		{
			"...123.456...",
			10,
			[]string{"456"},
		},
		{
			"...123.456...",
			11,
			[]string{},
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Input %s, Position %d", tc.inputLine, tc.position), func(t *testing.T) {
			res := ScanLineForNumbers([]rune(tc.inputLine), tc.position)
			assert.Equal(t, tc.want, res)
		})
	}
}

func TestScanGearRatios(t *testing.T) {
	type test struct {
		name      string
		inputFile string
		want      int
	}

	tests := []test{
		{
			"test 1",
			"testdata/1.txt",
			467835,
		},
		{
			"test 2",
			"testdata/2.txt",
			6756,
		},
		{
			"test 3",
			"testdata/3.txt",
			6756,
		}, {
			"test 4",
			"testdata/4.txt",
			2,
		},
		{
			"test 5",
			"testdata/5.txt",
			442,
		},
		{
			"test 6",
			"testdata/6.txt",
			999,
		},
		{
			"test 7",
			"testdata/7.txt",
			999,
		},
		{
			"test 8",
			"testdata/8.txt",
			999,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lines, err := util.ReadFileToStringSlice(tc.inputFile)
			if err != nil {
				t.Fatalf("Could not read inputfile, %v", err)
			}
			res := ScanGearRatios(lines)
			assert.Equal(t, tc.want, res)
		})
	}
}
