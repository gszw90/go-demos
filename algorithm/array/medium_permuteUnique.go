package array

import "slices"

// 47. 全排列 II
// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
// 示例 1：
//
// 输入：nums = [1,1,2]
// 输出：
// [[1,1,2],
//  [1,2,1],
//  [2,1,1]]
// 示例 2：
//
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//

func permuteUnique(nums []int) [][]int {
	var res [][]int
	l := len(nums)
	if l <= 1 {
		return [][]int{nums}
	}

	// 排序
	slices.Sort(nums)

	path := make([]int, l)
	// 记录路径上的元素是否已经被使用过
	onPath := make([]bool, l)

	var dfs func(i int)
	dfs = func(i int) {
		if i == l {
			res = append(res, append([]int{}, path...))
			return
		}
		for index, used := range onPath {
			if used {
				continue
			}
			// 去重,当前数字与上一个数字相同，且上一个数字没有被使用过
			if index > 0 && nums[index] == nums[index-1] && !onPath[index-1] {
				continue
			}
			// 标记当前元素为已使用
			path[i] = nums[index]
			onPath[index] = true
			dfs(i + 1)
			// 回溯，标记当前元素为未使用
			onPath[index] = false
		}
	}

	dfs(0)

	return res
}
