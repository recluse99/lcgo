/*
 * @lc app=leetcode.cn id=724 lang=golang
 *
 * [724] 寻找数组的中心下标
 */
package lcgo

// @lc code=start
func pivotIndex(nums []int) int {

	l := len(nums)

	if l == 1 {
		return 0
	}

	sum := sum724(nums)

	var leftSum int
	for i := 0; i < len(nums); i++ {
		// 防止   6   7     but 6==13/2=6   ——> return i
		if float64(leftSum) == float64(sum-nums[i])/2 {
			return i
		}

		// 后加
		leftSum += nums[i]

	}

	return -1
}

func sum724(nums []int) int {
	var sum int

	for _, value := range nums {
		sum += value
	}

	return sum
}

// @lc code=end
