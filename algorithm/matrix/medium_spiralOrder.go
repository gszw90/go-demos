package matrix

// 54.螺旋矩阵
// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
// 示例一
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]
// 示例二
// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// 输出：[1,2,3,4,8,12,11,10,9,5,6,7]

func spiralOrder(matrix [][]int) []int {
	// 没有数据
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	// 行与列
	rows, columns := len(matrix), len(matrix[0])
	// 记录是否访问过
	visited := make([][]bool, rows)
	for j := 0; j < rows; j++ {
		visited[j] = make([]bool, columns)
	}
	res := make([]int, rows*columns)
	total := rows * columns
	// 四个方向:右下左上
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	// 当前方向
	directionIndex := 0

	// 当前坐标
	row, column := 0, 0
	for i := 0; i < total; i++ {
		res[i] = matrix[row][column]
		visited[row][column] = true
		// 计算下一个坐标
		nextRow, nextColumn := row+directions[directionIndex][0], column+directions[directionIndex][1]
		// 判断是否越界或者已经访问过
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
			// 转向
			directionIndex = (directionIndex + 1) % 4
		}
		// 更新坐标
		row += directions[directionIndex][0]
		column += directions[directionIndex][1]
	}
	return res
}
