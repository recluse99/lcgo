/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 */
package lcgo

// @lc code=start
func findLengthOfLCIS(nums []int) int {

	l := len(nums)

	if l == 1 {
		return 1
	}

	firstIndex := 0
	secondIndex := 1

	maxLenght := 0

	for ; secondIndex < l && firstIndex < secondIndex; maxLenght++ {

		if nums[secondIndex-1] < nums[secondIndex] {
			maxLenght = max674(maxLenght, secondIndex-firstIndex+1)
		} else {
			firstIndex = secondIndex
		}
	}

	return maxLenght
}

func max674(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// @lc code=end
