package string

// 3442. 奇偶频次间的最大差值 I
// 给你一个由小写英文字母组成的字符串 s 。
// 请你找出字符串中两个字符 a1 和 a2 的出现频次之间的 最大 差值 diff = a1 - a2，这两个字符需要满足：
//
// a1 在字符串中出现 奇数次 。
// a2 在字符串中出现 偶数次 。
// 返回 最大 差值。
// 实例1
// 输入：s = "aaaaabbc"
// 输出：3
// 解释：a1 = 'a' 出现 5 次，a2 = 'b' 出现 2 次，两者之间的差值为 5 - 2 = 3 。
// 示例2
// 输入：s = "abcabcab"
// 输出：1
// 解释：a1 = 'a' 出现 3 次，a2 = 'c' 出现 2 次，两者之间的差值为 3 - 2 = 1 。

func maxDifference(s string) int {
	m := make(map[rune]int)

	for _, v := range s {
		m[v]++
	}
	// 最大奇数
	maxOdd := 0
	// 最小偶数
	minEven := len(s)

	for _, v := range m {
		if v%2 == 1 {
			maxOdd = max(maxOdd, v)
		} else {
			minEven = min(minEven, v)
		}
	}

	return maxOdd - minEven

}
