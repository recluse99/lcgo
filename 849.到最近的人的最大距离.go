/*
 * @lc app=leetcode.cn id=849 lang=golang
 *
 * [849] 到最近的人的最大距离
 */
package lcgo

// @lc code=start
func maxDistToClosest(seats []int) int {

	l := len(seats)

	if l == 2 {
		return 1
	}
	var result int
	var leftLen, rightLen int
	//中间值存放
	var temp int
	for i := 0; i < l; i++ {
		if seats[i] == 1 {
			continue
			// seats[i]==0
		} else {
			if i == 0 {
				temp = i
				for seats[temp] == 0 {
					temp++
				}
				rightLen = temp - i
				result = max849(result, rightLen)
			} else if i == l-1 {
				temp = i
				for seats[temp] == 0 {
					temp--
				}
				leftLen = i - temp
				result = max849(result, leftLen)
			} else {
				temp = i
				for temp < l && seats[temp] == 0 {
					temp++
				}
				rightLen = temp - i

				temp = i
				for temp >= 0 && seats[temp] == 0 {
					temp--
				}
				leftLen = i - temp

				temp = min849(leftLen, rightLen)

				result = max849(temp, result)

			}

		}


	}

	return result

}

func min849(a, b int) int {

	if a >= b {
		return b
	} else {
		return a
	}

}

func max849(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// @lc code=end
