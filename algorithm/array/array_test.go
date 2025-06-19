package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected int
	}{
		{
			name:     "empty input",
			words:    []string{},
			expected: 0,
		},
		{
			name:     "single palindrome word",
			words:    []string{"aa"},
			expected: 2,
		},
		{
			name:     "single non-palindrome word",
			words:    []string{"ab"},
			expected: 0,
		},
		{
			name:     "multiple palindrome words with odd count",
			words:    []string{"aa", "aa", "aa"},
			expected: 6,
		},
		{
			name:     "multiple palindrome words with even count",
			words:    []string{"aa", "aa", "aa", "aa"},
			expected: 8,
		},
		{
			name:     "pair of reverse words",
			words:    []string{"ab", "ba", "ab", "ba"},
			expected: 8,
		},
		{
			name:     "mixed palindrome and non-palindrome words",
			words:    []string{"aa", "ab", "ba", "aa"},
			expected: 8,
		},
		{
			name:     "multiple pairs with one palindrome middle",
			words:    []string{"ab", "ba", "cd", "dc", "ee"},
			expected: 10,
		},
		{
			name:     "all unique non-palindrome words",
			words:    []string{"ab", "cd", "ef"},
			expected: 0,
		},
		{
			name:     "multiple palindrome words with one odd count",
			words:    []string{"aa", "aa", "bb", "bb", "bb"},
			expected: 10,
		},
		{
			name:     "complex case with multiple pairs and middle",
			words:    []string{"ab", "ba", "cd", "dc", "ee", "ff", "ff"},
			expected: 14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, longestPalindrome(tt.words))
		})
	}
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:     "empty input",
			input:    []string{},
			expected: [][]string{},
		},
		{
			name:     "single word",
			input:    []string{"hello"},
			expected: [][]string{{"hello"}},
		},
		{
			name:     "multiple anagrams",
			input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name:     "no anagrams",
			input:    []string{"abc", "def", "ghi"},
			expected: [][]string{{"abc"}, {"def"}, {"ghi"}},
		},
		{
			name:     "case with duplicates",
			input:    []string{"listen", "silent", "listen"},
			expected: [][]string{{"listen", "silent", "listen"}},
		},
		{
			name:     "mixed case and anagrams",
			input:    []string{"rat", "art", "tar", "car"},
			expected: [][]string{{"rat", "art", "tar"}, {"car"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := groupAnagrams(tt.input)
			assert.ElementsMatch(t, tt.expected, actual)
		})
	}
}

func TestFindEvenNumbers(t *testing.T) {
	tests := []struct {
		name     string
		digits   []int
		expected []int
	}{
		{
			name:     "empty input",
			digits:   []int{},
			expected: []int{},
		},
		{
			name:     "less than 3 digits",
			digits:   []int{1, 2},
			expected: []int{},
		},
		{
			name:     "no possible even numbers",
			digits:   []int{1, 3, 5},
			expected: []int{},
		},
		{
			name:     "single valid even number",
			digits:   []int{2, 1, 3},
			expected: []int{132, 312},
		},
		{
			name:     "multiple valid even numbers",
			digits:   []int{2, 1, 3, 4},
			expected: []int{124, 132, 134, 142, 214, 234, 312, 314, 324, 342, 412, 432},
		},
		{
			name:     "duplicate digits",
			digits:   []int{2, 2, 8, 8},
			expected: []int{228, 282, 288, 822, 828, 882},
		},
		{
			name:     "all zeros",
			digits:   []int{0, 0, 0},
			expected: []int{},
		},
		{
			name:     "leading zeros not allowed",
			digits:   []int{0, 1, 2},
			expected: []int{102, 120, 210},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := findEvenNumbers(tt.digits)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "empty array",
			nums: []int{},
			want: [][]int{{}},
		},
		{
			name: "single element",
			nums: []int{1},
			want: [][]int{{1}},
		},
		{
			name: "two elements",
			nums: []int{1, 2},
			want: [][]int{
				{1, 2},
				{2, 1},
			},
		},
		{
			name: "three elements",
			nums: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.nums)
			assert.ElementsMatch(t, tt.want, got, "permute() = %v, want %v", got, tt.want)
		})
	}
}

func TestPermuteUnique(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "empty array",
			nums:     []int{},
			expected: [][]int{{}},
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: [][]int{{1}},
		},
		{
			name:     "two unique elements",
			nums:     []int{1, 2},
			expected: [][]int{{1, 2}, {2, 1}},
		},
		{
			name:     "two duplicate elements",
			nums:     []int{1, 1},
			expected: [][]int{{1, 1}},
		},
		{
			name:     "three elements with duplicates",
			nums:     []int{1, 1, 2},
			expected: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
		{
			name: "four elements with duplicates",
			nums: []int{1, 2, 2, 3},
			expected: [][]int{
				{1, 2, 2, 3}, {1, 2, 3, 2}, {1, 3, 2, 2},
				{2, 1, 2, 3}, {2, 1, 3, 2}, {2, 2, 1, 3},
				{2, 2, 3, 1}, {2, 3, 1, 2}, {2, 3, 2, 1},
				{3, 1, 2, 2}, {3, 2, 1, 2}, {3, 2, 2, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := permuteUnique(tt.nums)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}

func TestMaximumTripletValue(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int64
	}{
		{
			name:     "positive numbers with valid triplet",
			nums:     []int{12, 6, 1, 2, 7},
			expected: 77,
		},
		{
			name:     "all increasing numbers",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "all decreasing numbers",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 4,
		},
		{
			name:     "less than 3 elements",
			nums:     []int{1, 2},
			expected: 0,
		},
		{
			name:     "exactly 3 elements with valid triplet",
			nums:     []int{10, 5, 8},
			expected: 40,
		},
		{
			name:     "exactly 3 elements with invalid triplet",
			nums:     []int{5, 6, 7},
			expected: 0,
		},
		{
			name:     "negative numbers with valid triplet",
			nums:     []int{-5, -6, -1, -2, -7},
			expected: 5,
		},
		{
			name:     "mixed positive and negative numbers",
			nums:     []int{-1, 0, 1, -2, 3},
			expected: 9,
		},
		{
			name:     "all zeros",
			nums:     []int{0, 0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "single peak in middle",
			nums:     []int{2, 4, 1, 3, 5},
			expected: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, maximumTripletValue(tt.nums))
		})
	}
}

func TestMaximumTripletValue2(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int64
	}{
		{
			name:     "positive numbers with valid triplet",
			nums:     []int{12, 6, 1, 2, 7},
			expected: 77,
		},
		{
			name:     "all increasing numbers",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "all decreasing numbers",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 4,
		},
		{
			name:     "less than 3 elements",
			nums:     []int{1, 2},
			expected: 0,
		},
		{
			name:     "exactly 3 elements with valid triplet",
			nums:     []int{10, 5, 8},
			expected: 40,
		},
		{
			name:     "exactly 3 elements with invalid triplet",
			nums:     []int{5, 6, 7},
			expected: 0,
		},
		{
			name:     "negative numbers with valid triplet",
			nums:     []int{-5, -6, -1, -2, -7},
			expected: 5,
		},
		{
			name:     "mixed positive and negative numbers",
			nums:     []int{-1, 0, 1, -2, 3},
			expected: 9,
		},
		{
			name:     "all zeros",
			nums:     []int{0, 0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "single peak in middle",
			nums:     []int{2, 4, 1, 3, 5},
			expected: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, maximumTripletValue2(tt.nums))
		})
	}
}

func TestMinimumOperations(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int64
	}{
		{
			name:     "empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "no duplicates",
			nums:     []int{1, 2, 3, 4},
			expected: 0,
		},
		{
			name:     "duplicate at beginning",
			nums:     []int{1, 1, 2, 3},
			expected: 1, // i=1, i/3+1=0+1=1
		},
		{
			name:     "duplicate at middle",
			nums:     []int{1, 2, 3, 2, 4},
			expected: 1, // i=3, i/3+1=1+1=2
		},
		{
			name:     "duplicate at end",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1},
			expected: 1, // i=9, i/3+1=3+1=4
		},
		{
			name:     "multiple duplicates, return first found",
			nums:     []int{1, 2, 3, 4, 5, 2, 6, 7, 3},
			expected: 1, // i=5 is found before i=8
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumOperations(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMiniOperations(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "empty array",
			nums:     []int{},
			k:        5,
			expected: 0,
		},
		{
			name:     "all elements equal to k",
			nums:     []int{5, 5, 5},
			k:        5,
			expected: 0,
		},
		{
			name:     "all elements greater than k",
			nums:     []int{6, 7, 8},
			k:        5,
			expected: 3,
		},
		{
			name:     "some elements greater than k",
			nums:     []int{3, 5, 6, 7, 5, 8},
			k:        5,
			expected: -1,
		},
		{
			name:     "mixed elements with duplicates",
			nums:     []int{6, 7, 6, 8, 9, 8},
			k:        5,
			expected: 4,
		},
		{
			name:     "single element equal to k",
			nums:     []int{5},
			k:        5,
			expected: 0,
		},
		{
			name:     "single element greater than k",
			nums:     []int{6},
			k:        5,
			expected: 1,
		},
		{
			name:     "single element less than k",
			nums:     []int{4},
			k:        5,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := miniOperations(tt.nums, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCountBadPairs(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int64
	}{
		{
			name:     "empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: 0,
		},
		{
			name:     "all good pairs",
			nums:     []int{1, 2, 3, 4},
			expected: 0,
		},
		{
			name:     "all bad pairs",
			nums:     []int{4, 3, 2, 1},
			expected: 6, // C(4,2) = 6
		},
		{
			name:     "mixed pairs",
			nums:     []int{1, 2, 4, 4},
			expected: 3,
		},
		{
			name:     "large numbers",
			nums:     []int{100, 200, 300, 400},
			expected: 6,
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -2, -3, -4},
			expected: 6,
		},
		{
			name:     "mixed positive and negative",
			nums:     []int{-1, 0, 1, 2},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := countBadPairs(tt.nums)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestCountBadPairs2(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int64
	}{
		{
			name:     "empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: 0,
		},
		{
			name:     "all good pairs",
			nums:     []int{1, 2, 3, 4},
			expected: 0,
		},
		{
			name:     "all bad pairs",
			nums:     []int{4, 3, 2, 1},
			expected: 6, // C(4,2) = 6
		},
		{
			name:     "mixed pairs",
			nums:     []int{1, 2, 4, 4},
			expected: 3,
		},
		{
			name:     "large numbers",
			nums:     []int{100, 200, 300, 400},
			expected: 6,
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -2, -3, -4},
			expected: 6,
		},
		{
			name:     "mixed positive and negative",
			nums:     []int{-1, 0, 1, 2},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := countBadPairs2(tt.nums)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestCountSubarrays(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int64
	}{
		{
			name:     "empty array",
			nums:     []int{},
			k:        1,
			expected: 0,
		},
		{
			name:     "k larger than array length",
			nums:     []int{1, 2, 3},
			k:        4,
			expected: 0,
		},
		{
			name:     "single element array with k=1",
			nums:     []int{5},
			k:        1,
			expected: 1,
		},
		{
			name:     "all elements same and max",
			nums:     []int{3, 3, 3},
			k:        2,
			expected: 3,
		},
		{
			name:     "multiple max elements",
			nums:     []int{1, 3, 2, 3, 3},
			k:        2,
			expected: 6,
		},
		{
			name:     "no subarrays meet condition",
			nums:     []int{1, 2, 1, 2},
			k:        3,
			expected: 0,
		},
		{
			name:     "complex case 1",
			nums:     []int{1, 4, 2, 4, 3, 4},
			k:        2,
			expected: 10,
		},
		{
			name:     "complex case 2",
			nums:     []int{5, 1, 5, 2, 5, 3, 5},
			k:        3,
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := countSubarrays(tt.nums, tt.k)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_countCompleteSubarrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 3, 1, 2, 2},
			},
			want: 4,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{5, 5, 5, 5},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countCompleteSubarrays(tt.args.nums), "countCompleteSubarrays(%v)", tt.args.nums)
		})
	}
}

func TestMinZeroArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		queries  [][]int
		expected int
	}{
		{
			name:     "already zero sum",
			nums:     []int{0, 0, 0},
			queries:  [][]int{{0, 1, 1}},
			expected: 0,
		},
		{
			name:     "single query makes sum zero",
			nums:     []int{1, 2, 3},
			queries:  [][]int{{0, 3, 6}},
			expected: 1,
		},
		{
			name:     "multiple queries needed",
			nums:     []int{2, 0, 2},
			queries:  [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}},
			expected: 2,
		},
		{
			name:     "never reaches zero",
			nums:     []int{5, 5, 5},
			queries:  [][]int{{0, 1, 1}, {1, 2, 1}},
			expected: -1,
		},
		{
			name:     "partial range in query",
			nums:     []int{3, 1, 4, 2},
			queries:  [][]int{{1, 3, 3}, {0, 4, 5}},
			expected: 2,
		},
		{
			name:     "empty nums",
			nums:     []int{},
			queries:  [][]int{{0, 1, 1}},
			expected: 0,
		},
		{
			name:     "empty queries",
			nums:     []int{1, 2, 3},
			queries:  [][]int{},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := minZeroArray(tt.nums, tt.queries)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "all positive numbers",
			nums:     []int{1, 2, 3, 4},
			expected: 10,
		},
		{
			name:     "all negative numbers",
			nums:     []int{-1, -2, -3, -4},
			expected: -1,
		},
		{
			name:     "mixed positive and negative numbers",
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		{
			name:     "single element",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "with zeros",
			nums:     []int{0, -1, 0, 2, 0, -3},
			expected: 2,
		},
		{
			name:     "large numbers",
			nums:     []int{1000000, -1, 1000000},
			expected: 1999999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxSubArray(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCanJump_SingleElementArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Single element array with 0",
			nums:     []int{0},
			expected: true,
		},
		{
			name:     "Direct jump to last position from first",
			nums:     []int{5, 0, 0, 0, 0},
			expected: true,
		},
		{
			name:     "Cannot reach the end",
			nums:     []int{0, 1, 0, 0},
			expected: false,
		},
		{
			name:     "one",
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			name:     "one 2",
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canJump(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}
func TestLexicalOrder_NumberEndingWith9(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected []int
	}{
		{
			name:     "Number ending with 9",
			input:    19,
			expected: []int{1, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "Single transition from 9",
			input:    9,
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "Two-digit numbers",
			input:    15,
			expected: []int{1, 10, 11, 12, 13, 14, 15, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "Random middle number",
			input:    42,
			expected: []int{1, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 2, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 3, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 4, 40, 41, 42, 5, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lexicalOrder(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name:     "Non-overlapping intervals",
			input:    [][]int{{1, 3}, {4, 6}, {8, 10}},
			expected: [][]int{{1, 3}, {4, 6}, {8, 10}},
		},
		{
			name:     "Multiple overlapping intervals",
			input:    [][]int{{1, 3}, {2, 6}, {8, 10}},
			expected: [][]int{{1, 6}, {8, 10}},
		},
		{
			name:     "example 1",
			input:    [][]int{{1, 4}, {4, 5}},
			expected: [][]int{{1, 5}},
		},
		{
			name:     "Fully Contained Interval",
			input:    [][]int{{1, 6}, {2, 4}},
			expected: [][]int{{1, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := merge(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaxAdjacentDistance(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Basic Functionality with Increasing Sequence",
			input:    []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "example 1",
			input:    []int{1, 2, 4},
			expected: 3,
		},
		{
			name:     "example 2",
			input:    []int{-5, -10, -5},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxAdjacentDistance(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMinimizeMax(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		p        int
		expected int
	}{
		{
			name:     "Basic functionality with small input",
			nums:     []int{1, 2, 3, 4},
			p:        1,
			expected: 1,
		},
		{
			name:     "example 1",
			nums:     []int{10, 1, 2, 7, 1, 3},
			p:        2,
			expected: 1,
		},
		{
			name:     "example 2",
			nums:     []int{4, 2, 1, 2},
			p:        1,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimizeMax(tt.nums, tt.p)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaximumDifference(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Basic Increasing Sequence",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "example 1",
			nums:     []int{7, 1, 5, 4},
			expected: 4,
		},
		{
			name:     "example 2",
			nums:     []int{9, 4, 3, 2},
			expected: -1,
		},
		{
			name:     "example 3",
			nums:     []int{999, 997, 980, 976, 948, 940, 938, 928, 924, 917, 907, 907, 881, 878, 864, 862, 859, 857, 848, 840, 824, 824, 824, 805, 802, 798, 788, 777, 775, 766, 755, 748, 735, 732, 727, 705, 700, 697, 693, 679, 676, 644, 634, 624, 599, 596, 588, 583, 562, 558, 553, 539, 537, 536, 509, 491, 485, 483, 454, 449, 438, 425, 403, 368, 345, 327, 287, 285, 270, 263, 255, 248, 235, 234, 224, 221, 201, 189, 187, 183, 179, 168, 155, 153, 150, 144, 107, 102, 102, 87, 80, 57, 55, 49, 48, 45, 26, 26, 23, 15},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximumDifference(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		expected    [][]int
	}{
		{
			name:        "Empty intervals list",
			intervals:   [][]int{},
			newInterval: []int{1, 3},
			expected:    [][]int{{1, 3}},
		},
		{
			name:        "example 1",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			expected:    [][]int{{1, 5}, {6, 9}},
		},
		{
			name:        "example 2",
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			expected:    [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name:        "example 3",
			intervals:   [][]int{{1, 5}},
			newInterval: []int{2, 3},
			expected:    [][]int{{1, 5}},
		},
		{
			name:        "example 4",
			intervals:   [][]int{{1, 5}},
			newInterval: []int{2, 7},
			expected:    [][]int{{1, 7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := insert(tt.intervals, tt.newInterval)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestInsert2(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		expected    [][]int
	}{
		{
			name:        "Empty intervals list",
			intervals:   [][]int{},
			newInterval: []int{1, 3},
			expected:    [][]int{{1, 3}},
		},
		{
			name:        "example 1",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			expected:    [][]int{{1, 5}, {6, 9}},
		},
		{
			name:        "example 2",
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			expected:    [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name:        "example 3",
			intervals:   [][]int{{1, 5}},
			newInterval: []int{2, 3},
			expected:    [][]int{{1, 5}},
		},
		{
			name:        "example 4",
			intervals:   [][]int{{1, 5}},
			newInterval: []int{2, 7},
			expected:    [][]int{{1, 7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := insert2(tt.intervals, tt.newInterval)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPartitionArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{3, 6, 1, 2, 5},
			k:        2,
			expected: 2,
		},
		{
			name:     "example 2",
			nums:     []int{1, 2, 3},
			k:        1,
			expected: 2,
		},
		{
			name:     "example 3",
			nums:     []int{2, 2, 4, 5},
			k:        0,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := partitionArray(tt.nums, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}
