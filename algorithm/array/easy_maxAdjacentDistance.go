package array

// 3423. 循环数组中相邻元素的最大差值
// 给你一个 循环 数组 nums ，请你找出相邻元素之间的 最大 绝对差值。
// 注意：一个循环数组中，第一个元素和最后一个元素是相邻的。
// 示例一
// 输入：nums = [1,2,4]
// 输出：3
// 解释：由于 nums 是循环的，nums[0] 和 nums[2] 是相邻的，它们之间的绝对差值是最大值 |4 - 1| = 3 。
// 示例二
// 输入：nums = [-5,-10,-5]
// 输出：5
// 解释：nums[0] 和 nums[1] 是相邻的，它们之间的绝对差值是最大值 |-5 - (-10)| = 5 。

func maxAdjacentDistance(nums []int) int {
	maxDis := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		var tmpDis int
		if i == 0 {
			tmpDis = nums[i] - nums[l-1]
		} else {
			tmpDis = nums[i] - nums[i-1]
		}

		if tmpDis < 0 {
			tmpDis *= -1
		}
		maxDis = max(maxDis, tmpDis)
	}
	return maxDis
}
