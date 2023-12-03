package main

import (
	"github.com/stretchr/testify/assert"
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
