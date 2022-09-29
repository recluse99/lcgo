/*
 * @lc app=leetcode.cn id=1005 lang=golang
 *
 * [1005] K 次取反后最大化的数组和
 */
package lcgo

import "sort"

// @lc code=start
func largestSumAfterKNegations(nums []int, k int) int {

	//进行升序
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		//处理负值
		if nums[i] < 0 && k > 0 {
			nums[i] = -nums[i]
			k--
		}
		if k == 0 {
			return sum(nums)
		}

		/* if i<len(nums)-1&&nums[i+1] > 0  {
			break
		} */
	}

	//到这说明k>0
	sort.Ints(nums)
	if k%2 == 0 {
		return sum(nums)
	} else {
		nums[0] = -nums[0]
		return sum(nums)
	}

}

func sum(nums []int) int {

	var sum int

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum

}

// @lc code=end
