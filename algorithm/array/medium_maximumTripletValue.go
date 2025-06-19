package array

/*
 * @desc:给你一个下标从 0 开始的整数数组 nums 。
 *       请你从所有满足 i < j < k 的下标三元组 (i, j, k) 中，找出并返回下标三元组的最大值。如果所有满足条件的三元组的值都是负数，则返回 0 。
 *       下标三元组 (i, j, k) 的值等于 (nums[i] - nums[j]) * nums[k] 。
 * @example:
 *          输入：nums = [12,6,1,2,7]
 *          输出：77
 *          解释：下标三元组 (0, 2, 4) 的值是 (nums[0] - nums[2]) * nums[4] = 77 。
 *          可以证明 77 是满足题目要求的最大值。
 * @note: 3 <= nums.length <= 105
 *        1 <= nums[i] <= 106
 */

/*
 * @solution: 固定第二个数字j，取前j-1个数字中最大的数字，取后j+1个数字中最大的数字，计算最大值
 *            时间复杂度：O(n^2)
 *            空间复杂度：O(1)
 */
func maximumTripletValue(nums []int) int64 {
	l := len(nums)
	if l < 3 {
		return 0
	}
	leftMax := make([]int, l)
	rightMax := make([]int, l)
	leftMax[0] = nums[0]
	rightMax[l-1] = nums[l-1]
	for i := 1; i < l-1; i++ {
		leftMax[i] = max(leftMax[i-1], nums[i-1])
	}
	for i := l - 2; i >= 1; i-- {
		rightMax[i] = max(rightMax[i+1], nums[i+1])
	}
	maxResult := 0
	for i := 1; i < l-1; i++ {
		maxResult = max(maxResult, (leftMax[i]-nums[i])*rightMax[i])
	}
	return int64(maxResult)
}

/*
 * @solution: 固定第三个数字k，取前k-1个数字中差值最大的值与现在相乘，计算最大值
 */

func maximumTripletValue2(nums []int) int64 {
	l := len(nums)
	if l < 3 {
		return 0
	}
	maxResult := 0
	maxNum := max(nums[0], nums[1])
	maxDiff := nums[0] - nums[1]
	for i := 2; i < l; i++ {
		maxResult = max(maxResult, maxDiff*nums[i])
		maxDiff = max(maxDiff, maxNum-nums[i])
		maxNum = max(maxNum, nums[i])
	}

	return int64(maxResult)
}
