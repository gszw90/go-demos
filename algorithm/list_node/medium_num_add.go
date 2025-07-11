package list_node

/**
 * @desc: 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
 *        请你将两个数相加，并以相同形式返回一个表示和的链表。
 *        你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 * @example1: 输入：l1 = [2,4,3], l2 = [5,6,4]
 *           输出：[7,0,8]
 *           解释：342 + 465 = 807.
 * @example2: 输入：l1 = [0], l2 = [0]
 *           输出：[0]
 * @example3: 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
 *           输出：[8,9,9,9,0,0,0,1]
 * @note: 每个链表中的节点数在范围 [1, 100] 内
 *        0 <= Node.val <= 9
 *        题目数据保证列表表示的数字不含前导零
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummy.Next
}
