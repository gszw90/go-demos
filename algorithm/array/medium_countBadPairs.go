package array

// 2364. 统计坏数对的数目
// 给你一个下标从 0 开始的整数数组 nums 。如果 i < j 且 j - i != nums[j] - nums[i] ，那么我们称 (i, j) 是一个 坏数对 。
//
// 请你返回 nums 中 坏数对 的总数目。
//
// 示例 1：
//
// 输入：nums = [4,1,3,3]
// 输出：5
// 解释：数对 (0, 1) 是坏数对，因为 1 - 0 != 1 - 4 。
// 数对 (0, 2) 是坏数对，因为 2 - 0 != 3 - 4, 2 != -1 。
// 数对 (0, 3) 是坏数对，因为 3 - 0!= 3 - 4, 3!= -1 。
// 数对 (1, 2) 是坏数对，因为 2 - 1!= 3 - 1, 1!= 2 。
// 数对 (2, 3) 是坏数对，因为 3 - 2!= 3 - 3, 1!= 0 。
// 总共有 5 个坏数对，所以我们返回 5 。
// 示例 2：
//
// 输入：nums = [1,2,3,4,5]
// 输出：0
// 解释：没有坏数对。
//
// 提示：
//
// 1 <= nums.length <= 105
// 1 <= nums[i] <= 109

func countBadPairs(nums []int) int64 {
	// 思路：
	// 1. 遍历数组，计算每个元素与其他元素的差的绝对值
	// 2. 如果差的绝对值不等于元素的下标差，则为坏数对
	// 3. 统计坏数对的数量
	// 4. 返回坏数对的数量
	var count int64
	l := len(nums)
	if l < 2 {
		return count
	}
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if j-i != nums[j]-nums[i] {
				count++
			}
		}
	}

	return count
}

func countBadPairs2(nums []int) int64 {
	// 思路：
	// 转换为公式为:i-nums[i]!=j-nums[j]

	l := len(nums)
	if l < 2 {
		return 0
	}
	// 表示位置i处左边i-nums[i]的数量
	m := make(map[int]int)
	// 好数对的最大数量
	// 通过遍历获取好对数的数量，减去之后就只剩下坏对数的数量
	total := l * (l - 1) / 2

	for i := 0; i < l; i++ {
		key := nums[i] - i
		// 减去好对数的数量
		total -= m[key]
		m[key]++
	}

	return int64(total)
}
