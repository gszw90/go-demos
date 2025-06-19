package string

import "strconv"

// 2843. 统计对称整数的数目
// 给你两个正整数 low 和 high 。
//
// 对于一个由 2 * n 位数字组成的整数 x ，如果其前 n 位数字之和与后 n 位数字之和相等，则认为这个数字是一个对称整数。
//
// 返回在 [low, high] 范围内的 对称整数的数目 。
//
// 示例 1：
// 输入：low = 1, high = 100
// 输出：9
// 解释：在 1 到 100 范围内共有 9 个对称整数：11、22、33、44、55、66、77、88 和 99 。
//
// 示例 2：
// 输入：low = 1200, high = 1230
// 输出：4
// 解释：在 1200 到 1230 范围内共有 4 个对称整数：1203、1212、1221 和 1230 。
//
// 提示：
// 1 <= low <= high <= 10000
//
// 思路：从 low 到 high 遍历，判断每个数是否为对称整数，如果是则计数器加 1
//
// 注意：1. 对称整数的前 n 位数字之和与后 n 位数字之和相等
//      2. 对称整数的位数为偶数
//      3. 对称整数的前 n 位数字之和与后 n 位数字之和相等的条件是：前 n 位数字之和等于后 n 位数字之和

func countSymmetricIntegers(low int, high int) int {
	count := 0
	for i := low; i <= high; i++ {
		if isSymmetric(i) {
			count++
		}
	}
	return count
}

func countSymmetricIntegers2(low int, high int) int {
	count := 0
	for i := low; i <= high; i++ {
		if i > 10 && i < 100 {
			if i%11 == 0 {
				count++
			}
		} else if i > 1000 && i < 10000 {
			leftSum := i/1000 + (i%1000)/100
			rightSum := (i % 100 / 10) + i%10
			if leftSum == rightSum {
				count++
			}
		}

	}
	return count
}

func isSymmetric(num int) bool {
	s := strconv.Itoa(num)
	n := len(s)
	if n%2 != 0 {
		return false
	}
	leftSum, rightSum := 0, 0
	for i := 0; i < n/2; i++ {
		leftSum += int(s[i] - '0')
		rightSum += int(s[n-i-1] - '0')
	}
	return leftSum == rightSum
}
