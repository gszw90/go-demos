package sort

// select sort
// 每次从未排序的部分选出最小（或最大）的元素，然后与未排序部分的第一个元素交换位置，如此反复直到整个数组排序完成
// 时间复杂度：O(n^2)
func SelectSort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	for i := 0; i < l; i++ {
		minIndex := i
		for j := i + 1; j < l; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}
	return nums
}
