/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */
package lcgo

import (
	"math"
)

// @lc code=start
func thirdMax(nums []int) int {

	//比 nums的最小值还小  ，可以更好的判断three是否改变
	var first int = math.MinInt64
	var second int = math.MinInt64
	var three int = math.MinInt64

	for _, value := range nums {

		if value > first {
			three = second
			second = first
			first = value
		} else if value > second && value < first {
			three = second
			second = value
		} else if value > three && value < second {
			three = value
		}

	}
	if three == math.MinInt64 {
		return first
	} else {
		return three
	}

}

// @lc code=end
