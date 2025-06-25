package list_node

// 61. 旋转链表
// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
//
// 示例 1：
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[4,5,1,2,3]
//
// 示例 2：
// 输入：head = [0,1,2], k = 4
// 输出：[2,0,1]

func rotateRight(head *ListNode, k int) *ListNode {
	// 1. 通过遍历获取链表长度
	// 2. 转换成换装链表
	// 3. 计算环状链表断开的位置，断开链表
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	current := head
	// 循环完成后current指向最后一个节点
	for current.Next != nil {
		n++
		current = current.Next
	}

	// 计算断开处的位置
	pos := n - k%n

	if pos == n {
		return head
	}

	// 构建环装列表
	current.Next = head

	for pos > 0 {
		pos--
		current = current.Next
	}
	// 断开节点
	res := current.Next
	current.Next = nil
	return res
}
