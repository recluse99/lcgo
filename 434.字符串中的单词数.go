/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 */
package lcgo

import "strings"

// @lc code=start
func countSegments(s string) int {

	// 去不去掉都可
	result := strings.Trim(s, " ")

	var count int

	l := len(result)

	for i := 0; i < l; i++ {
		if result[i] != ' ' {
			count++

			// 防止越界
			for i < l && result[i] != ' ' {
				i++
			}
		}
	}

	return count
}

// @lc code=end
