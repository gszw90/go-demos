package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	k := 3
	res := subArray(s, k)
	fmt.Println(res)
	// originalNode := createNode([]int{1, 2, 3, 4, 5})
	// printNode(originalNode)
	//
	// revNode := reverseNode(originalNode)
	// printNode(revNode)

}

// [2, 5, 8, 10, 12]
// 1. 创建一个单链表
// 2. 翻转一个单链表

type node struct {
	val  int
	next *node
}

func createNode(s []int) *node {
	dummy := &node{}
	current := dummy
	for _, v := range s {
		current.next = &node{
			val: v,
		}
		current = current.next
	}
	return dummy.next
}

func printNode(n *node) {
	current := n
	for current != nil {
		fmt.Print(current.val, " ")
		current = current.next
	}
	fmt.Println()
}

func reverseNode(n *node) *node {
	dummy := &node{}
	current := n
	for current != nil {
		next := current.next
		current.next = dummy.next
		dummy.next = current
		current = next
	}

	return dummy.next
}

func createNode1(s []int) {
	head := &node{}
	current := head
	// 创建节点
	for _, v := range s {
		current.next = &node{
			val: v,
		}
		current = current.next
	}
	// 反转节点
	revNode := &node{}
	current = head.next
	for current != nil {
		next := current.next
		current.next = revNode.next
		revNode.next = current
		current = next
	}
	fmt.Println("finish reverse nodes")
	// 打印原始节点
	current = head.next
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
	fmt.Println("======================")
	current = revNode.next
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}

}

func subArray(s []int, k int) [][]int {

	res := [][]int{}
	var dfs func(prefix []int, i int)
	dfs = func(prefix []int, i int) {

		if len(prefix) == k {
			res = append(res, prefix)
			return
		}
		if i >= len(s) {
			return
		}
		tmp := make([]int, len(prefix))
		copy(tmp, prefix[:])
		// 使用当前位置的值
		tmp = append(tmp, s[i])
		dfs(tmp, i+1)
		// 不使用当前位置的值
		dfs(prefix, i+1)
	}
	dfs([]int{}, 0)
	return res
}

func set(s []int, k int) [][]int {
	res := [][]int{}
	var dfs func(prefix []int, i int)
	dfs = func(prefix []int, i int) {
		if len(prefix) == k {
			res = append(res, prefix)
			return
		}

		tmp := []int{}
		copy(tmp, prefix[:])
		// 使用当前位置的值
		tmp = append(tmp, s[i])
		dfs(tmp, i+1)
		// 不使用当前位置的值
		dfs(prefix, i+1)
	}

	dfs([]int{}, 0)
	return res

}
