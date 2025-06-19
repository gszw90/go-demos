package array

// 2962. 统计最大元素出现至少 K 次的子数组
// 给你一个下标从 0 开始的整数数组 nums 和一个正整数 k 。
//
// 请你统计并返回 nums 中 最大 元素 至少出现 k 次的 非空 子数组的数目。
//
// 子数组是数组中一个 连续 元素序列。
//
// 示例 1：
//
// 输入：nums = [1,3,2,3,3], k = 2
// 输出：6
// 输出： 包含元素 3 至少 2 次的子数组为：[1,3,2,3]、[1,3,2,3,3]、[3,2,3]、[3,2,3,3]、[2,3,3] 和 [3,3]
// 示例 2：
//
// 输入：nums = [1,4,2,1], k = 3
// 输出：0
// 解释：
// 不存在包含元素 4 至少 3 次的子数组。

func countSubarrays(nums []int, k int) int64 {
	// 思路：滑动窗口
	// 1. 遍历数组，找到最大的元素
	// 2. 从左到右遍历数组，找到以i为结尾的子数组中，最大元素为maxNum的子数组的数量
	// 3. 从左到右遍历数组，找到以i为开头的子数组中，最大元素为maxNum的子数组的数量
	// 4. 遍历数组，找到以i为结尾的子数组中，最大元素为maxNum的子数组的数量

	// 遍历数组，找到最大的元素
	l := len(nums)
	if l < k {
		return 0
	}

	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	// 从left到当前位置最大元素个数
	cnt := 0
	res := 0
	left := 0
	for _, num := range nums {
		if num == maxNum {
			cnt++
		}
		// 获取从left到i的子数组中，最短的子数组位置
		for cnt >= k {
			if nums[left] == maxNum {
				cnt--
			}
			left++
		}
		res += left
	}

	return int64(res)
}
