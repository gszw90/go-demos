package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "empty array",
			nums:     []int{},
			target:   5,
			expected: -1,
		},
		{
			name:     "single element array - target found",
			nums:     []int{5},
			target:   5,
			expected: 0,
		},
		{
			name:     "single element array - target not found",
			nums:     []int{3},
			target:   5,
			expected: -1,
		},
		{
			name:     "target at beginning",
			nums:     []int{1, 3, 5, 7, 9},
			target:   1,
			expected: 0,
		},
		{
			name:     "target at end",
			nums:     []int{1, 3, 5, 7, 9},
			target:   9,
			expected: 4,
		},
		{
			name:     "target in middle",
			nums:     []int{1, 3, 5, 7, 9},
			target:   5,
			expected: 2,
		},
		{
			name:     "target in left half",
			nums:     []int{1, 3, 5, 7, 9},
			target:   3,
			expected: 1,
		},
		{
			name:     "target in right half",
			nums:     []int{1, 3, 5, 7, 9},
			target:   7,
			expected: 3,
		},
		{
			name:     "target not in array",
			nums:     []int{1, 3, 5, 7, 9},
			target:   4,
			expected: -1,
		},
		{
			name:     "all elements same - target found",
			nums:     []int{2, 2, 2, 2, 2},
			target:   2,
			expected: 2, // any index with value 2 is correct
		},
		{
			name:     "all elements same - target not found",
			nums:     []int{2, 2, 2, 2, 2},
			target:   3,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinarySearch(tt.nums, tt.target)
			if tt.name == "all elements same - target found" {
				// For this special case, just verify that we found some occurrence
				if tt.expected != -1 {
					assert.Equal(t, tt.target, tt.nums[result])
				} else {
					assert.Equal(t, tt.expected, result)
				}
			} else {
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
