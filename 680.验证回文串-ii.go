/*
 * @lc app=leetcode.cn id=680 lang=golang
 *
 * [680] 验证回文串 II
 */

package lcgo

// @lc code=start
func validPalindrome(s string) bool {

	return twoPointAddFirst(s) || twoPointReduceEnd(s)

}

func twoPointAddFirst(s string) bool {
	l := len(s)
	firstIndex := 0
	endIndex := l - 1
	var count int

	for firstIndex < endIndex {

		if s[firstIndex] == s[endIndex] {
			firstIndex++
			endIndex--
		} else {

			if count == 1 {
				return false
			}

			if s[firstIndex+1] == s[endIndex] {
				firstIndex++
			} else {
				return false
			}
			// 删除一字符
			count++
		}

	}

	return true

}
func twoPointReduceEnd(s string) bool {
	l := len(s)
	firstIndex := 0
	endIndex := l - 1
	var count int

	for firstIndex < endIndex {

		if s[firstIndex] == s[endIndex] {
			firstIndex++
			endIndex--
		} else {

			if count == 1 {
				return false
			}

			if s[firstIndex] == s[endIndex-1] {
				endIndex--
			} else {
				return false
			}
			// 删除一字符
			count++
		}

	}

	return true

}

// @lc code=end
