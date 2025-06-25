package matrix

// 59. 螺旋矩阵 II
// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
//
// 示例 1：
// 输入：n = 3
// 输出：[[1,2,3],[8,9,4],[7,6,5]]
//
// 示例 2：
// 输入：n = 1
// 输出：[[1]]

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	// 方向
	directs := [][]int{
		{0, 1},  // 右
		{1, 0},  // 下
		{0, -1}, // 左
		{-1, 0}, // 上
	}
	direstIdx := 0
	row, col := 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		// 超出边界或者已经被占用
		nextRow := row + directs[direstIdx][0]
		nextCol := col + directs[direstIdx][1]
		if nextRow >= n || nextRow < 0 || nextCol >= n || nextCol < 0 || matrix[nextRow][nextCol] > 0 {
			direstIdx = (direstIdx + 1) % 4
		}
		row += directs[direstIdx][0]
		col += directs[direstIdx][1]

	}

	return matrix
}
