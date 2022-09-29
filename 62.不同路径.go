/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */
package lcgo

// @lc code=start
func uniquePaths(m int, n int) int {

	//当只有一行或一列时只有一条路线
	if m == 1 || n == 1 {
		return 1
	}

	//go 中数组容量只能用常量创建
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	//m>1
	dp[0][0] = 1
	dp[0][1] = 1
	dp[1][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if i == 0 && j == 0 {
				continue
			}

			if i == 0 {
				dp[i][j] = dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}

		}
	}

	return dp[m-1][n-1]
}

// @lc code=end
