/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start
package lcgo

func candy(ratings []int) int {

	if len(ratings) == 1 {
		return 1
	}

	children := make([]int, len(ratings))

	//每人一颗
	for i := 0; i < len(children); i++ {
		children[i] = 1
	}

	//var total int = len(ratings)
	//从左向右   right children > left children
	for i := 1; i < len(ratings); i++ {

		if ratings[i] > ratings[i-1] {

			children[i] = children[i-1] + 1
		}

	}

	//从右向左  left children > right children
	for i := len(ratings) - 1; i > 0; i-- {
		//左边高分孩子得到更多的糖果
		if ratings[i-1] > ratings[i] {
			//至少保证左边高分孩子糖果比右边低分孩子多一颗糖果
			if children[i-1] <= children[i] {
				children[i-1] = children[i] + 1
			}
		}
	}

	return numsOneThreeFive(children)
}

func numsOneThreeFive(a []int) int {
	nums := 0
	for i := 0; i < len(a); i++ {
		nums += a[i]
	}
	return nums
}

// @lc code=end
