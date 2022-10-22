/*
 * @lc app=leetcode.cn id=665 lang=golang
 *
 * [665] 非递减数列
 */
package lcgo

// @lc code=start
func checkPossibility(nums []int) bool {

	l := len(nums)

	if l <= 2 {
		return true
	}
	count := 0

	for i := 0; i < l-1; i++ {

		if nums[i] > nums[i+1] {

			count++

			if count > 1 {
				return false
			}
			// 还没有想明白
			if i > 0 && nums[i-1] > nums[i+1] {
				nums[i+1] = nums[i]
			}

		}

	}

	return true
}

// @lc code=end
