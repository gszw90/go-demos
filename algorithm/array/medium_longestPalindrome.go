package array

import "strings"

// 2131.连接两字母单词得到的最长回文串
// 给你一个字符串数组 words 。words 中每个元素都是一个包含 两个 小写英文字母的单词。
//
// 请你从 words 中选择一些元素并按 任意顺序 连接它们，并得到一个 尽可能长的回文串 。每个元素 至多 只能使用一次。
//
// 请你返回你能得到的最长回文串的 长度 。如果没办法得到任何一个回文串，请你返回 0 。
//
// 回文串 指的是从前往后和从后往前读一样的字符串。
// 示例 1：
// 输入：words = ["lc","cl","gg"]
// 输出：6
// 解释：一个最长的回文串为 "lc" + "gg" + "cl" = "lcggcl" ，长度为 6 。
// "clgglc" 是另一个可以得到的最长回文串。
// 示例 2：
// 输入：words = ["ab","ty","yt","lc","cl","ab"]
// 输出：8
// 解释：最长回文串是 "ty" + "lc" + "cl" + "yt" = "tylcclyt" ，长度为 8 。
// "lcyttycl" 是另一个可以得到的最长回文串。
// 示例 3：
// 输入：words = ["cc","ll","xx"]
// 输出：2

func longestPalindrome(words []string) int {
	// 记录每个单词出现的次数
	m := make(map[string]int)
	for _, v := range words {
		m[v]++
	}
	l := 0
	hasMid := false
	for w, cnt := range m {
		// 获取反转后的单词
		rev := string([]byte{w[1], w[0]})
		// 如果反转后的单词和原单词相同，那么可以选择两个相同的单词作为回文串的中间部分
		// 如果反转后的单词和原单词不同，那么可以选择两个不同的单词作为回文串的两边部分
		if rev == w {
			// 如果单词出现的次数是奇数，那么可以选择一个作为回文串的中间部分
			if cnt%2 == 1 {
				hasMid = true
			}
			l += cnt / 2 * 4
		} else if strings.Compare(w, rev) > 0 {
			// 使用compare函数比较两个字符串的大小，避免重复计算
			l += min(cnt, m[rev]) * 4
		}

	}
	if hasMid {
		return l + 2
	}
	return l
}
