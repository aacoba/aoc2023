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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := GetFirstAndLastDigit(tc.input)
			if !tc.wantErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
			assert.Equal(t, tc.expected, result)
		})
	}
}
