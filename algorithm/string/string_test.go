package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthAfterTransformations(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        int
		expected int
	}{
		{
			name:     "empty string",
			s:        "",
			t:        5,
			expected: 0,
		},
		{
			name:     "abcyy string",
			s:        "abcyy",
			t:        2,
			expected: 7,
		},
		{
			name:     "no transformations",
			s:        "abc",
			t:        0,
			expected: 3,
		},
		{
			name:     "single transformation",
			s:        "a",
			t:        1,
			expected: 1, // a -> b
		},
		{
			name:     "single transformation with z",
			s:        "z",
			t:        1,
			expected: 2, // z -> a
		},
		{
			name:     "multiple transformations",
			s:        "ab",
			t:        2,
			expected: 2, // ab -> bc -> cd
		},
		{
			name:     "transformations with wrap around",
			s:        "yz",
			t:        2,
			expected: 4, // yz -> zab -> abbc
		},
		{
			name:     "multiple transformations with duplicate letters",
			s:        "aab",
			t:        3,
			expected: 3, // aab -> bbc -> ccd -> dde
		},
		{
			name:     "large number of transformations",
			s:        "abc",
			t:        26,
			expected: 6, // should wrap around to original letters
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := lengthAfterTransformations(tt.s, tt.t)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMinimumCost(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64
	}{
		{
			name:     "example1",
			input:    "0011",
			expected: 2,
		},
		{
			name:     "example2",
			input:    "010101",
			expected: 9,
		},
		{
			name:     "all zeros",
			input:    "0000",
			expected: 0,
		},
		{
			name:     "all ones",
			input:    "1111",
			expected: 0,
		},
		{
			name:     "single character",
			input:    "0",
			expected: 0,
		},
		{
			name:     "alternating 01",
			input:    "0101",
			expected: 4, // 2 + 2
		},
		{
			name:     "alternating 10",
			input:    "1010",
			expected: 4, // 2 + 2
		},
		{
			name:     "long string with one transition",
			input:    "00000000001",
			expected: 1,
		},
		{
			name:     "multiple transitions",
			input:    "00110011",
			expected: 8,
		},
		{
			name:     "complex pattern",
			input:    "001011100",
			expected: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumCost(tt.input)
			if result != tt.expected {
				t.Errorf("minimumCost(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "single character",
			input:    "a",
			expected: 1,
		},
		{
			name:     "all unique characters",
			input:    "abcde",
			expected: 5,
		},
		{
			name:     "repeating characters",
			input:    "aaaaa",
			expected: 1,
		},
		{
			name:     "substring in middle",
			input:    "abcabcbb",
			expected: 3,
		},
		{
			name:     "substring at start",
			input:    "bbbbb",
			expected: 1,
		},
		{
			name:     "substring at end",
			input:    "pwwkew",
			expected: 3,
		},
		{
			name:     "mixed characters",
			input:    "aab",
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lengthOfLongestSubstring(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaxDifference(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "example1",
			input:    "aaaaabbc",
			expected: 3,
		},
		{
			name:     "example2",
			input:    "abcabcab",
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxDifference(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
func TestMinMaxDifference(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Single digit number",
			input:    5,
			expected: 9,
		},
		{
			name:     "example 1",
			input:    11891,
			expected: 99009,
		},
		{
			name:     "example 2",
			input:    90,
			expected: 99,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minMaxDifference(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaxDiff(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "All digits are 9",
			input:    9999,
			expected: 8888,
		},
		{
			name:     "example 1",
			input:    555,
			expected: 888,
		},
		{
			name:     "example 2",
			input:    123456,
			expected: 820000,
		},
		{
			name:     "example 3",
			input:    10000,
			expected: 80000,
		},
		{
			name:     "example 4",
			input:    111,
			expected: 888,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxDiff(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
