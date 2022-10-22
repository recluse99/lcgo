/*
 * @lc app=leetcode.cn id=830 lang=golang
 *
 * [830] 较大分组的位置
 */
package lcgo

// @lc code=start
func largeGroupPositions(s string) [][]int {

	l := len(s)
	if l < 3 {
		return [][]int{}
	}

	result := [][]int{}

	for i := 0; i < l-2; i++ {
		// s[i] is 较大分组
		if s[i] == s[i+1] && s[i+1] == s[i+2] {
			result = append(result, []int{i, i + 2})
			temp := s[i]
			//i移动到最后一个字符的index
			for i < l-1 && s[i+1] == temp { //短路与   前面false  后面不再执行
				i++
			}
			result[len(result)-1][1] = i

		}

	}

	return result
}

// @lc code=end
