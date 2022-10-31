/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */
package lcgo

// @lc code=start
func isIsomorphic(s string, t string) bool {
	ls := len(s)
	lt := len(t)

	if ls != lt {
		return false
	}

	return isIsomorphicRecur(s, t) && isIsomorphicRecur(t, s)

}

func isIsomorphicRecur(s string, t string) bool {

	ls := len(s)

	rMap := map[byte]byte{}
	for i := 0; i < ls; i++ {
		if _, ok := rMap[s[i]]; !ok {
			rMap[s[i]] = t[i]
			//存在
		} else {
			if rMap[s[i]] != t[i] {
				return false
			}

		}
	}

	return true

}

// @lc code=end
