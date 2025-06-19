package array

// 57. 插入区间
// 给你一个 无重叠的 ，按照区间起始端点排序的区间列表 intervals，
// 其中 intervals[i] = [starti, endi] 表示第 i 个区间的开始和结束，并且 intervals 按照 starti 升序排列。
// 同样给定一个区间 newInterval = [start, end] 表示另一个区间的开始和结束。
// 在 intervals 中插入区间 newInterval，使得 intervals 依然按照 starti 升序排列，且区间之间不重叠（如果有必要的话，可以合并区间）。
// 返回插入之后的 intervals。
// 注意 你不需要原地修改 intervals。你可以创建一个新数组然后返回它。
//
// 示例 1：
// 输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
// 输出：[[1,5],[6,9]]
//
// 示例 2：
// 输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
// 输出：[[1,2],[3,10],[12,16]]
// 解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。

func insert(intervals [][]int, newInterval []int) [][]int {
	l := len(intervals)
	res := make([][]int, 0, l)
	if l == 0 {
		res = append(res, newInterval)
		return res
	}
	// 插入元素在左边
	if newInterval[1] < intervals[0][0] {
		res = append(res, newInterval)
		res = append(res, intervals...)
		return res
	}
	// 插入元素在右边
	if newInterval[0] > intervals[l-1][1] {
		return append(intervals, newInterval)
	}

	current := newInterval
	for key, interval := range intervals {
		interval := interval
		// 左边重叠区域
		if interval[1] < current[0] {
			res = append(res, interval)
			continue
		}
		// 右边无重叠区域
		if interval[0] > current[1] {
			res = append(res, current)
			res = append(res, intervals[key:]...)
			break
		}
		// 重叠区域
		// 插入元素在某个元素的的区间范围内
		if interval[0] <= newInterval[0] && interval[1] >= newInterval[1] {
			res = append(res, intervals[key:]...)
			break
		}

		// 左重叠
		if interval[0] <= current[0] {
			current[0] = interval[0]
		}
		// 右重叠
		if interval[1] >= current[1] {
			current[1] = interval[1]
		}
		// 最后一个
		if key == l-1 {
			res = append(res, current)
		}

	}

	return res
}

func insert2(intervals [][]int, newInterval []int) [][]int {
	l := len(intervals)
	res := make([][]int, 0, l)
	if l == 0 {
		res = append(res, newInterval)
		return res
	}
	left, right := newInterval[0], newInterval[1]
	// 插入元素是否已经合并到结果集
	isMerged := false

	for key, interval := range intervals {
		if interval[1] < left {
			// 当前元素在插入元素的左边
			res = append(res, interval)

		} else if interval[0] > right {
			// 当前元素在插入元素的右边

			// 判断新的元素是否被插入到结果集
			if !isMerged {
				isMerged = true
				res = append(res, []int{left, right})
			}
			// 将后面的元素都插入结果集
			res = append(res, intervals[key:]...)
			break

		} else {
			// 重叠区域
			left = min(interval[0], left)
			right = max(interval[1], right)
		}
	}

	if !isMerged {
		res = append(res, []int{left, right})
	}

	return res
}
