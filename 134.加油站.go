/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
package lcgo

func canCompleteCircuit(gas []int, cost []int) int {

	//出发的地址
	var index int

	var temp int
	//如果总消耗大于总油量，则不可能环形一周
	if sumOneThreeFour(cost) > sumOneThreeFour(gas) {
		return -1
	}

	//每个加油站剩余的油
	rest := make([]int, len(gas))

	/* 从零开始累计rest中的数，当累计的和小于0时，
	说明再区间[0,index]区间都不能为出发点 */

	for i := 0; i < len(gas); i++ {
		
		rest[i] = gas[i] - cost[i]

		temp += rest[i]

		if temp < 0 {
			index = i + 1
			//重置
			temp = 0
		}
	}

	return index

}

func sumOneThreeFour(a []int) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}

	return sum
}

// @lc code=end
