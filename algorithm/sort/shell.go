package sort

// shell sort
// 使用不同的步长将数组分成几个子数组，针对每个子数组进行插入排序，然后逐渐减小步长，如此反复直到整个数组排序完成。
// 时间复杂度：O(nlogn)
func ShellSort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	// 初始步长
	gap := l / 2
	for gap > 0 {
		// 对每个子数组进行插入排序
		for i := gap; i < l; i++ {
			temp := nums[i]
			// 从后往前遍历，找到第一个小于当前元素的位置
			for j := i - gap; j >= 0; j -= gap {
				if nums[j] < temp {
					break
				}
				if nums[j] > nums[j+gap] {
					nums[j], nums[j+gap] = nums[j+gap], nums[j]
				}
			}
		}
		// 减小步长
		gap /= 2
	}
	return nums
}
