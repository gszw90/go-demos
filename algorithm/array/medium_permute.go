package array

// 46. 全排列
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
// 示例 1：
//
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 示例 2：
//
// 输入：nums = [0,1]
// 输出：[[0,1],[1,0]]
// 示例 3：
//
// 输入：nums = [1]
// 输出：[[1]]
//

func permute(nums []int) [][]int {
	var res [][]int
	l := len(nums)

	if l <= 1 {
		return [][]int{nums}
	}

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
			// 如果当前元素未被被使用过
			if !used {
				// 标记当前元素为已使用
				path[i] = nums[index]
				onPath[index] = true
				dfs(i + 1)
				// 回溯，标记当前元素为未使用
				onPath[index] = false
			}
		}
	}
	dfs(0)
	return res
}
