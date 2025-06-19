package sort

// insert sort
// 直接插入排序是一种简单的插入排序法，其基本思想是：从第二个元素开始，将每个元素插入到已排序的数组中的适当位置，直到整个数组排序完成。
// 时间复杂度：O(n^2)
func InsertSort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	for i := 1; i < l; i++ {
		temp := nums[i]
		// 从后往前遍历，找到第一个小于当前元素的位置
		for j := i - 1; j >= 0; j-- {
			if nums[j] < temp {
				break
			}
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
