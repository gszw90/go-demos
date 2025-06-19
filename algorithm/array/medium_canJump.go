package array

// 55. 跳跃游戏
// 给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
// 实例1
// 输入：nums = [2,3,1,1,4]
// 输出：true
// 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
// 示例2
// 输入：nums = [3,2,1,0,4]
// 输出：false

func canJump(nums []int) bool {
	l := len(nums)
	// 记录当前能跳到的最远位置
	maxPos := 0
	for i := 0; i < l; i++ {

		// 如果当前位置已经大于能跳到的最远位置，说明无法到达最后一个位置
		if i > maxPos {
			return false
		}
		// 能够跳到最后一个位置
		if i+nums[i] >= l-1 {
			return true
		}
		// 更新能跳到的最远位置
		maxPos = max(maxPos, i+nums[i])
	}

	return true
}
