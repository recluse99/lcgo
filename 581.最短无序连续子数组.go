/*
 * @lc app=leetcode.cn id=581 lang=golang
 *
 * [581] 最短无序连续子数组
 */
package lcgo

// @lc code=start

func findUnsortedSubarray(nums []int) int {

	l := len(nums)

	if l == 1 {
		return 0
	}

	var minIndex = 0
	var maxIndex = l - 1
	var maxValue = nums[0]
	var minValue = nums[l-1]

	// 从左向右 找到逆区间的右边界
	for i := 0; i < l; i++ {
		maxValue = max581(maxValue, nums[i])

		if nums[i] < maxValue {
			minIndex = i
		}

	}

	// 从右到左 找到逆区间的左边界
	for i := l - 1; i >= 0; i-- {
		minValue = min581(minValue, nums[i])
		// 不能取相等   取相等会导致右边界可能向右多移
		// 左移同理
		if nums[i] > minValue {
			maxIndex = i
		}
	}

	if minIndex <= maxIndex {
		return 0
	} else {
		return minIndex - maxIndex + 1
	}

}

// 比较大小
func min581(a int, b int) int {

	if a > b {
		return b
	} else {
		return a
	}
}

func max581(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
