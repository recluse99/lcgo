/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */
package lcgo

// @lc code=start

func minCostClimbingStairs(cost []int) int {

	var step = len(cost)

	dp := make([]int, step+1)
	//台阶阶数最低是2 ,所以dp[0]为1台阶的情况不存在，赋值1
	dp[0] = 0
	//无三元运算符
	if cost[0] < cost[1] {
		dp[1] = cost[0]
	} else {
		dp[1] = cost[1]
	}

	//当cost数组为长度为2时， 有两台阶
	if step == 2 {
		return dp[step-1]
	}

	//step> 2

	for i := 2; i < step; i++ {
		//此时有两种情况
		if dp[i-1]+cost[i] <= dp[i-2]+cost[i-1] {
			dp[i] = dp[i-1] + cost[i]
		} else {
			dp[i] = dp[i-2] + cost[i-1]
		}

	}

	return dp[step-1]

}

// @lc code=end
