/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球
 */
package lcgo

import "sort"

// @lc code=start
func findMinArrowShots(points [][]int) int {

	if len(points) == 0 {
		return 0
	}
	//以左边界排序   从小到大  如果第一个数相等,那么判断第二个数
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		} else {
			return points[i][0] < points[j][0]
		}

	})

	var count int

	var j int

	for i := 0; i < len(points); i++ {

		count++

		temp := points[i][1]

		for j = i + 1; j < len(points); j++ {
			if temp >= points[j][0] && temp <= points[j][1] || points[j][1] < temp {
				if j == len(points)-1 {
					return count
				}
				continue
			} else {

				break
			}
		}
		i = j - 1

	}

	return count
}

// @lc code=end
