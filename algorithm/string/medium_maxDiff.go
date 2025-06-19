package string

import (
	"strconv"
	"strings"
)

// 1432. 改变一个整数能得到的最大差值
// 给你一个整数 num 。你可以对它进行以下步骤共计 两次：
// 选择一个数字 x (0 <= x <= 9).
// 选择另一个数字 y (0 <= y <= 9) 。数字 y 可以等于 x 。
// 将 num 中所有出现 x 的数位都用 y 替换。
// 令两次对 num 的操作得到的结果分别为 a 和 b 。
// 请你返回 a 和 b 的 最大差值 。
// 注意，新的整数（a 或 b）必须不能 含有前导 0，并且 非 0。
//
// 示例 1：
// 输入：num = 555
// 输出：888
// 解释：第一次选择 x = 5 且 y = 9 ，并把得到的新数字保存在 a 中。
// 第二次选择 x = 5 且 y = 1 ，并把得到的新数字保存在 b 中。
// 现在，我们有 a = 999 和 b = 111 ，最大差值为 888
//
// 示例 2：
// 输入：num = 9
// 输出：8
// 解释：第一次选择 x = 9 且 y = 9 ，并把得到的新数字保存在 a 中。
// 第二次选择 x = 9 且 y = 1 ，并把得到的新数字保存在 b 中。
// 现在，我们有 a = 9 和 b = 1 ，最大差值为 8
//
// 示例 3：
// 输入：num = 123456
// 输出：820000

func maxDiff(num int) int {
	// 最大的数，找到第一位非9的数替换为9
	// 最小的数：第一位不是1替换成1，剩余非0的数替换成0
	s := strconv.Itoa(num)
	l := len(s)
	maxPos := -1
	minPos := -1
	maxStr, minStr := s, s
	isAllOne := true
	for i := 0; i < l; i++ {
		if maxPos > -1 && minPos > -1 {
			break
		}
		if s[i] > '1' {
			isAllOne = false
		}
		// 获取最大值
		if s[i] != '9' && maxPos == -1 {
			maxPos = i
			maxStr = strings.ReplaceAll(maxStr, string(s[i]), "9")
		}
		// 前面所有的值都为1
		if isAllOne && s[i] == '1' {
			continue
		}
		// 获取最小值
		if minPos == -1 {
			// 第一位不为0
			if i == 0 && s[i] > '1' {
				minPos = i
				minStr = strings.ReplaceAll(minStr, string(s[i]), "1")
			}
			// 其他位置上不为0
			if i > 0 && s[i] > '0' {
				minPos = i
				minStr = strings.ReplaceAll(minStr, string(s[i]), "0")
			}
		}
	}
	maxNum, _ := strconv.Atoi(maxStr)
	minNum, _ := strconv.Atoi(minStr)

	return maxNum - minNum

}
