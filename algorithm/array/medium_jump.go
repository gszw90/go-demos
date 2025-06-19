package array

// 45. 跳跃游戏 II
// 给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
// 假设你总是可以到达数组的最后一个位置。
//
// 示例 1:
// 输入: nums = [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//
// 示例 2:
// 输入: nums = [2,3,0,1,4]
// 输出: 2
//
// 提示:
// 1 <= nums.length <= 104
// 0 <= nums[i] <= 1000

func jump(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}
	// 当前位置到达的最远位置的距离
	maxPosition := l - 1
	steps := 0
	for maxPosition > 0 {
		// 找到当前位置到达的最远位置
		for i := 0; i < maxPosition; i++ {
			if i+nums[i] >= maxPosition {
				maxPosition = i
				steps++
				break
			}
		}
	}

	return steps

}
