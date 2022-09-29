/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */
package lcgo

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	//单行单列
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if m == 1 || n == 1 {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if obstacleGrid[i][j] == 1 {
					return 0
				}
			}
		}

		return 1
	}
	//终点被堵
	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	//m>1 n>1
	//创建dp切片
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	//给dp的开头赋值
	if obstacleGrid[0][0] != 1 {
		dp[0][0] = 1
	} else {
		dp[0][0] = 0
	}

	// if obstacleGrid[0][1] != 1 {
	// 	dp[0][1] = 1
	// } else {
	// 	dp[0][1] = 0
	// }

	// if obstacleGrid[1][0] != 1 {
	// 	dp[1][0] = 1
	// } else {
	// 	dp[1][0] = 0
	// }

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			
			if i == 0 && j == 0 {
				continue
			}

			if j == 0 {
				if obstacleGrid[i][j] != 1 {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = 0
				}

			} else if i == 0 {
				if obstacleGrid[i][j] != 1 {
					dp[i][j] = dp[i][j-1]
				} else {
					dp[i][j] = 0
				}

			} else {
				if obstacleGrid[i][j] != 1 {
					dp[i][j] = dp[i][j-1] + dp[i-1][j]
				} else {
					dp[i][j] = 0
				}

			}

		}

	}

	return dp[m-1][n-1]
}

// @lc code=end
