package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFirstAndLastDigit(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected int
		wantErr  bool
	}
	tests := []test{
		{
			name:     "Simple",
			input:    "123abc",
			expected: 13,
			wantErr:  false,
		},
		{
			name:     "no ints",
			input:    "abc",
			expected: 0, // doesnt care
			wantErr:  true,
		},
		{
			name:     "1 int",
			input:    "a1c",
			expected: 11,
			wantErr:  false,
		},
		{
			name:     "tokenization 1",
			input:    "two1nine",
			expected: 29,
			wantErr:  false,
		},
		{
			name:     "tokenization 2",
			input:    "eightwothree",
			expected: 83,
			wantErr:  false,
		},
		{"tokenization", "7pqrstsixteen", 76, false},
		{"tokenization", "abcone2threexyz", 13, false},
		{"tokenization", "xtwone3four", 24, false},
		{"tokenization", "4nineeightseven2", 42, false},
		{"tokenization", "zoneight234", 14, false},
		{"tokenization", "7pqrstsixteen", 76, false},
		{"tokenization", "gttsix2567", 67, false},
		{
			name:     "with gliberish",
			input:    "heksenkaastwo",
			expected: 22,
			wantErr:  false,
		},
		{"becasue fuck you thats why", "oneight", 18, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := FancyScannerGetFirstAndLastDigits(tc.input)
			if !tc.wantErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestScanStringForDigitTokens(t *testing.T) {
	type test struct {
		name    string
		input   string
		want    []string
		wantErr bool
	}

	tests := []test{
		{
			name:    "simple",
			input:   "123",
			want:    []string{"1", "2", "3"},
			wantErr: false,
		},
		{
			name:    "simple",
			input:   "onetwothree",
			want:    []string{"1", "2", "3"},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ScanStringForDigitTokens(tc.input)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tc.want, result)
		})
	}
}
