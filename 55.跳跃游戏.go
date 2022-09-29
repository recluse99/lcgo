/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */
package lcgo

// @lc code=start
func canJump(nums []int) bool {

	var cover int
	if len(nums) == 1 {
		return true
	}
	//覆盖数的初始化
	cover = nums[0]

	/*随着index的右移，可以覆盖的cover也在不断更新，
		当cover大于等于len(nums)-1时，返回true
	*/
	
	for i := 0; i <= cover; i++ {

		if i+nums[i] > cover {
			cover = i + nums[i]
		}

		if cover >= len(nums)-1 {
			return true
		}

	}

	return false
}

// @lc code=end
