package array

// 386. 字典序排数
// 给定一个整数 n, 返回从 1 到 n 的字典顺序。
// 例如，
// 给定 n =1 3，返回 [1,10,11,12,13,2,3,4,5,6,7,8,9] 。
// 请尽可能的优化算法的时间复杂度和空间复杂度。 输入的数据 n 小于等于 5,000,000。
// 示例 1:
// 输入: n = 13
// 输出: [1,10,11,12,13,2,3,4,5,6,7,8,9]

func lexicalOrder(n int) []int {
	res := make([]int, n)
	// 当前数字
	num := 1
	for i := 0; i < n; i++ {
		res[i] = num
		// 当前数字乘以10后，如果小于n，说明后面还有数字
		if num*10 <= n {
			num *= 10
		} else {
			// 如果当前数字各位为9或者本身等于n，说明后面没有数字了，返回想上找
			for num%10 == 9 || num == n {
				num /= 10
			}
			num++
		}
	}
	return res
}
