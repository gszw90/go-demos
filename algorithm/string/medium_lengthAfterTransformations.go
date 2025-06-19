package string

// 3335. 字符串转换后长度
// 给你一个字符串 s 和一个整数 t，表示要执行的 转换 次数。每次 转换 需要根据以下规则替换字符串 s 中的每个字符：
// 如果字符是 'z'，则将其替换为字符串 "ab"。
// 否则，将其替换为字母表中的下一个字符。例如，'a' 替换为 'b'，'b' 替换为 'c'，依此类推。
// 请你返回执行 t 次转换后得到的字符串。
// 示例 1：
// 输入：s = "abcyy", t = 2
// 输出：7
// 解释：执行 2 次转换后，字符串变为 "cdeabab"。

func lengthAfterTransformations(s string, t int) int {
	var mod = 1000000007
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[c-'a']++
	}
	// 遍历轮数
	for i := 0; i < t; i++ {
		// 暂时保存本轮的数据
		tmp := make([]int, 26)
		// 对a进行处理
		tmp[0] = cnt[25]
		// 对b进行处理
		tmp[1] = (cnt[0] + cnt[25]) % mod
		// 从c开始，对每个字母进行处理
		for j := 2; j < 26; j++ {
			tmp[j] = cnt[j-1]
		}
		// 将本轮的数据赋值给cnt
		cnt = tmp

	}
	var res int
	// 计算结果
	for _, v := range cnt {
		res = (res + v) % mod
	}
	return res
}
