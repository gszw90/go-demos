package array

import (
	"sort"
)

// 2616. 最小化数对的最大差值
// 给你一个下标从 0 开始的整数数组 nums 和一个整数 p 。
// 请你从 nums 中找到 p 个下标对，每个下标对对应数值取差值，你需要使得这 p 个差值的 最大值 最小。
// 同时，你需要确保每个下标在这 p 个下标对中最多出现一次。
// 对于一个下标对 i 和 j ，这一对的差值为 |nums[i] - nums[j]| ，其中 |x| 表示 x 的 绝对值 。
// 请你返回 p 个下标对对应数值 最大差值 的 最小值 。
//
// 示例 1：
// 输入：nums = [10,1,2,7,1,3], p = 2
// 输出：1
// 解释：第一个下标对选择 1 和 4 ，第二个下标对选择 2 和 5 。
// 最大差值为 max(|nums[1] - nums[4]|, |nums[2] - nums[5]|) = max(0, 1) = 1 。所以我们返回 1 。
//
// 示例 2：
// 输入：nums = [4,2,1,2], p = 1
// 输出：0
// 解释：选择下标 1 和 3 构成下标对。差值为 |2 - 2| = 0 ，这是最大差值的最小值。

func minimizeMax(nums []int, p int) int {
	// 对数组进行排序，要使得差值最小，必然是相邻的两个数的差值最小
	// 差值的范围为0到nums[len(nums)-1]-nums[0]
	// 为了快速找到这个最小值，可以使用二分法对插值范围进行查找
	sort.Ints(nums)
	l := len(nums)

	var check func(mx int) bool
	check = func(mx int) bool {
		cnt := 0
		for i := 0; i < l-1; i++ {
			// 相邻两个数差值小于max
			if nums[i+1]-nums[i] <= mx {
				cnt++
				i++
			}
		}
		return cnt >= p
	}
	// 差值范围
	left := 0
	right := nums[l-1] - nums[0]
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}
