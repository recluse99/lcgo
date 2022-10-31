/*
 * @lc app=leetcode.cn id=1518 lang=golang
 *
 * [1518] 换水问题
 */
package lcgo

// @lc code=start
func numWaterBottles(numBottles int, numExchange int) int {
	//喝水的瓶数
	var sum = 0
	// 空瓶
	var drink = 0
	for numBottles > 0 {
		numBottles--
		sum++
		drink++
		if drink == numExchange {
			numBottles++
			drink = 0
		}
	}

	return sum

}

// @lc code=end
