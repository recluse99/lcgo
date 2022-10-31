/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */
package lcgo

// @lc code=start
func maxArea(height []int) int {

	l := len(height)

	first := 0
	end := l - 1

	// because l >=2
	// 初始值   设为
	maxValue := 0

	for first < end {

		//value := min11(height[first], height[end]) * (end - first)
		// if value > maxValue {
		// 	maxValue = value
		// }
		// 高度小的一边移动，
		if height[first] <= height[end] {
			maxValue = max11(maxValue, height[first]*(end-first))
			first++
		} else {
			maxValue = max11(maxValue, height[end]*(end-first))
			end--
		}

	}

	return maxValue
}

func max11(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// @lc code=end
