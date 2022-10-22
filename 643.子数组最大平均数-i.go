/*
 * @lc app=leetcode.cn id=643 lang=golang
 *
 * [643] 子数组最大平均数 I
 */
package lcgo

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {

	l := len(nums)
	if l == k {
		return float64(sum643(nums)) / float64(k)
	}
	//当index为0时的sum
	maxValue := sum643(nums[0:k])
	temp := maxValue
	//从1开始遍历
	for i := 1; i <= l-k; i++ {
		
		temp = temp - nums[i-1] + nums[i+k-1]
		maxValue = max643(temp, maxValue)

	}

	return float64(maxValue) / float64(k)
}

func max643(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}

}

func sum643(a []int) int {

	sum := 0

	for _, value := range a {
		sum += value
	}

	return sum
}

// @lc code=end

//  过于浪费时间
// func findMaxAverage(nums []int, k int) float64 {

// 	l := len(nums)
// 	if l == k {
// 		return float64(sum643(nums)) / float64(k)
// 	}
// 	maxValue := math.MinInt32

// 	for i := 0; i <= l-k; i++ {

// 		maxValue = max643(sum643(nums[i:i+k]), maxValue)

// 	}

// 	return float64(maxValue) / float64(k)
// }

// func max643(a, b int) int {
// 	if a >= b {
// 		return a
// 	} else {
// 		return b
// 	}

// }

// func sum643(a []int) int {

// 	sum := 0

// 	for _, value := range a {
// 		sum += value
// 	}

// 	return sum
// }
