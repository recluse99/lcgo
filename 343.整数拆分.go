/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */
package lcgo

// @lc code=start
func integerBreak(n int) int {

	dp := make([]int, n+1)

	//因为n>=2
	dp[2] = 1

	// j =1  遍历到i-2
	//两种情况 ，j*(i-j)  和 j*dp[i-j]
	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], max(dp[i-j]*j, j*(i-j)))
		}
	}

	return dp[n]
}

func max(a int, b int) int {

	if a >= b {
		return a
	} else {
		return b
	}
}
// @lc code=end