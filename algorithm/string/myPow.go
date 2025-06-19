package string

// 50. Pow(x, n)
// 实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。
//
// 示例 1：
//
// 输入：x = 2.00000, n = 10
// 输出：1024.00000
// 示例 2：
//
// 输入：x = 2.10000, n = 3
// 输出：9.26100
// 示例 3：
//
// 输入：x = 2.00000, n = -2
// 输出：0.25000
// 解释：2-2 = 1/22 = 1/4 = 0.25

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	} else {
		return 1.0 / quickMul(x, -n)
	}
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	res := 1.0
	// 快速幂算法
	base := x
	for n > 0 {
		// 如果n是奇数，就把多出来的一个x乘到res中
		if n&1 == 1 {
			res *= base
		}
		// 把base平方，n除以2
		base *= base
		n >>= 1
	}
	return res
}
