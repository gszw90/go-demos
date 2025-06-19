package string

/*
 * @desc: 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
 * @example1: 输入: s = "abcabcbb"
 *            输出: 3
 *            解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 * @example2: 输入: s = "bbbbb"
 *            输出: 1
 *            解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 * @example3: 输入: s = "pwwkew"
 *            输出: 3
 *            解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * @note: 0 <= s.length <= 5 * 104
 *        s 由英文字母、数字、符号和空格组成
 *
 * @solution: 1. 遍历字符串，记录当前位置的字符和长度
 *            2. 如果当前位置的字符已经存在于map中，那么需要更新长度
 *            3. 如果当前位置的字符不存在于map中，那么需要更新长度
 *            4. 返回最大长度
 */

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l == 0 || l == 1 {
		return l
	}
	maxLength := 1
	left := 0
	right := 1
	m := make(map[byte]int)
	m[s[left]] = left
	for right < l {
		// 如果当前位置的字符已经存在于map中，那么需要更新长度
		if index, ok := m[s[right]]; ok && index >= left {
			left = index + 1
		}
		m[s[right]] = right
		right++
		if right-left > maxLength {
			maxLength = right - left
		}
	}
	return maxLength
}
