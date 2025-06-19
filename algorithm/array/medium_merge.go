package array

import "sort"

// 56. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
// 实例1
// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
// 输出：[[1,6],[8,10],[15,18]]
// 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
// 示例2
// 输入：intervals = [[1,4],[4,5]]
// 输出：[[1,5]]
// 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

func merge(intervals [][]int) [][]int {
	// 排序
	sort.Slice(intervals, func(i, j int) bool {

		return intervals[i][0] < intervals[j][0]

	})
	res := make([][]int, 0)
	// 合并区间
	current := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if current[1] >= intervals[i][0] {
			// 有重叠区域
			current[1] = max(current[1], intervals[i][1])
		} else {
			// 没有重叠区域
			res = append(res, current)
			current = intervals[i]
		}
	}
	// 最后一段
	res = append(res, current)

	return res
}
