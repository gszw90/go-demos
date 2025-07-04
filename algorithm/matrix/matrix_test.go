package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralOrder_EmptyMatrix(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected []int
	}{
		{
			name:     "Empty matrix (0 rows)",
			matrix:   [][]int{},
			expected: []int{},
		},
		{
			name:     "Single column matrix",
			matrix:   [][]int{{1}, {2}, {3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Rectangular Matrix Test (3x2)",
			matrix:   [][]int{{1, 2}, {3, 4}, {5, 6}},
			expected: []int{1, 2, 4, 6, 5, 3},
		},
		{
			name:     "3x3 square matrix",
			matrix:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := spiralOrder(tt.matrix)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGenerateMatrix(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected [][]int
	}{
		{
			name:     "Smallest valid input (1x1 matrix)",
			input:    1,
			expected: [][]int{{1}},
		},
		{
			name:     "n=2 (Small even size)",
			input:    2,
			expected: [][]int{{1, 2}, {4, 3}},
		},
		{
			name:  "Small odd size matrix (n=3)",
			input: 3,
			expected: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateMatrix(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
