/*
 * @lc app=leetcode.cn id=1513 lang=golang
 *
 * [1513] 仅含 1 的子串数
 */
package lcgo

// @lc code=start
func numSub(s string) int {

	l := len(s)
	var index, count, sum = 0, 0, 0

	for i := 0; i < l; i++ {
		if s[i] == '1' {
			index = i
			for i < l && s[i] == '1' {
				i++
				count = i - index
				sum += count
			}
		}
	}

	return sum % 1000000007
}

// @lc code=end

// 思路是对的。实现过程有点繁琐  // 优化 将 求和 加入到for循环中

// func numSub(s string) int {

// 	l := len(s)
// 	var index, count, sum = 0, 0, 0

// 	for i := 0; i < l; i++ {
// 		if s[i] == '1' {
// 			index = i
// 			for i < l && s[i] == '1' {
// 				i++
// 			}

// 			count = i - index

// 			sum += getCount(count)

// 		}
// 	}

// 	return sum % (10 ^ 9 + 7)
// }

// func getCount(count int) int {

// 	if count == 1 {
// 		return 1
// 	}
// 	return count + getCount(count-1)

// }
