package list_node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name     string
		l1       *ListNode
		l2       *ListNode
		expected *ListNode
	}{
		{
			name:     "Both lists are nil",
			l1:       nil,
			l2:       nil,
			expected: nil,
		},
		{
			name:     "One list is nil",
			l1:       &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			l2:       nil,
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
		{
			name:     "Both lists have single digits",
			l1:       &ListNode{Val: 2},
			l2:       &ListNode{Val: 3},
			expected: &ListNode{Val: 5},
		},
		{
			name:     "Lists with multiple digits",
			l1:       &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			l2:       &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			expected: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
		{
			name:     "Lists with carry over",
			l1:       &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}},
			l2:       &ListNode{Val: 1},
			expected: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}},
		},
		{
			name:     "Lists with different lengths",
			l1:       &ListNode{Val: 1, Next: &ListNode{Val: 8}},
			l2:       &ListNode{Val: 0},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 8}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addTwoNumbers(tt.l1, tt.l2)
			assert.True(t, compareLists(tt.expected, result))
		})
	}
}

// 添加自定义链表比较函数
func compareLists(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
