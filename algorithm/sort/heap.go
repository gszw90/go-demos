package sort

// heap sort
// 堆排序是一种基于二叉堆数据结构的排序算法，它的基本思想是：将待排序的数组构建成一个最大堆（或最小堆），然后将堆顶元素与堆底元素交换位置，
// 再将剩余的元素重新构建成一个最大堆（或最小堆），如此反复直到整个数组排序完成。
// 时间复杂度：O(nlogn)

func heapSort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	// 构建最大堆, 从最后一个非叶子节点开始
	for i := l/2 - 1; i >= 0; i-- {
		heapfy(nums, l, i)
	}
	// 依次将堆顶元素与堆底元素交换位置，再将剩余的元素重新构建成一个最大堆
	for i := l - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapfy(nums, i, 0)
	}
	return nums
}

func heapfy(nums []int, n, i int) {
	largestIndex := i // 假设当前节点为最大节点
	left := 2*i + 1   // 左子节点的索引
	right := 2*i + 2  // 右子节点的索引

	// 如果左子节点的值大于最大节点的值，则更新最大节点的索引
	if left < n && nums[left] > nums[largestIndex] {
		largestIndex = left
	}
	// 如果右子节点的值大于最大节点的值，则更新最大节点的索引
	if right < n && nums[right] > nums[largestIndex] {
		largestIndex = right
	}

	// 如果最大节点的索引不是当前节点的索引，则交换当前节点和最大节点的值
	if largestIndex != i {
		nums[i], nums[largestIndex] = nums[largestIndex], nums[i]

		// 递归地对最大节点进行堆化
		heapfy(nums, n, largestIndex)
	}
}
