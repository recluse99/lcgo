/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */
package lcgo

// @lc code=start
func firstUniqChar(s string) int {

	// 优化 ，使用切片存储字符

	l := len(s)

	rMap := map[byte]int{}

	for i := 0; i < l; i++ {
		rMap[s[i]]++
	}

	for i := 0; i < l; i++ {
		if rMap[s[i]] == 1 {
			return i
		}
	}

	return -1

}

// @lc code=end
