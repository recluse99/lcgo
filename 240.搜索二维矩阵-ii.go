/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */
package lcgo

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {

	m := len(matrix)
	n := len(matrix[0])
	// var i int
	// var j int

	// 从 右上 向左下
	for i, j := 0, n-1; i < m && j >= 0; {

		if target < matrix[i][j] {
			j--
		} else if target > matrix[i][j] {
			i++
		} else {
			return true
		}
	}

	return false

}

// @lc code=end
