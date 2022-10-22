/*
 * @lc app=leetcode.cn id=747 lang=golang
 *
 * [747] 至少是其他数字两倍的最大数
 */
package lcgo

// @lc code=start
func dominantIndex(nums []int) int {

	l := len(nums)

	if l == 1 {
		return 0
	}
	// 题目提示 nums[i]>=0
	var maxValue = 0
	var result = 0

	for index, value := range nums {
		if value > maxValue {
			maxValue = value
			result = index
		}
	}

	for index, value := range nums {
		//跳过最大值本身
		if index == result {
			continue
		}
		if maxValue < value*2 {
			return -1
		}
	}

	return result
}

// @lc code=end

//存在问题的解法   。。  无法处理在 除了index的其他数相等的情况

// func dominantIndex(nums []int) int {

// 	l := len(nums)

// 	if l == 1 {
// 		return 0
// 	}

// 	//切片的总和 和 0的个数
// 	sum, count := sum747(nums)

// 	for index, value := range nums {
// 		if value == 0 {
// 			if 0 > 2*(sum-value) {
// 				return index
// 			}
// 			//value!=0
// 		} else {
// 			if value*(l-1-count) > 2*(sum-value) {
// 				return index
// 			}
// 		}

// 	}

// 	return -1
// }

// func sum747(a []int) (int, int) {
// 	var sum int
// 	var count int
// 	for _, value := range a {
// 		if value == 0 {
// 			count++
// 		}
// 		sum += value
// 	}
// 	return sum, count
// }
