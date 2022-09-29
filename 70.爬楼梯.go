/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */
package lcgo

// @lc code=start
func climbStairs(n int) int {

	if n <= 2 {
		return n
	}
	//多申请一位，是index和n对应
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// @lc code=end
