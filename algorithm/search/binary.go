package search

// BinarySearch 二分查找
// 时间复杂度：O(logn)
func BinarySearch(nums []int, target int) int {
	l := len(nums)
	if l < 1 {
		return -1
	}
	left, right := 0, l-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
