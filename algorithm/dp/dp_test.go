package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostPoints(t *testing.T) {
	tests := []struct {
		name      string
		questions [][]int
		expected  int64
	}{
		{
			name:      "empty input",
			questions: [][]int{},
			expected:  0,
		},
		{
			name: "single question without skip",
			questions: [][]int{
				{5, 0},
			},
			expected: 5,
		},
		{
			name: "single question with skip",
			questions: [][]int{
				{5, 1},
			},
			expected: 5,
		},
		{
			name: "multiple questions with optimal path",
			questions: [][]int{
				{3, 2},
				{4, 3},
				{5, 0},
				{2, 1},
			},
			expected: 7,
		},
		{
			name: "all questions require skipping",
			questions: [][]int{
				{1, 3},
				{2, 2},
				{3, 1},
				{4, 0},
			},
			expected: 4,
		},
		{
			name: "large skip jumps over multiple questions",
			questions: [][]int{
				{10, 5},
				{1, 1},
				{1, 1},
				{1, 1},
				{1, 1},
				{1, 1},
				{5, 0},
			},
			expected: 15,
		},
		{
			name: "mixed value questions",
			questions: [][]int{
				{12, 4},
				{8, 1},
				{7, 0},
				{5, 2},
				{3, 1},
				{4, 0},
			},
			expected: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, mostPoints(tt.questions))
		})
	}
}
