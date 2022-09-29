/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */
package lcgo

// @lc code=start
func lemonadeChange(bills []int) bool {

	if bills[0] != 5 {
		return false
	}

	//存放零钱
	//money :=make(map[int]int,3)
	money := map[int]int{
		5:  0,
		10: 0,
		15: 0,
	}

	money[5] = 1

	for i := 1; i < len(bills); i++ {

		if bills[i] == 5 {
			money[5] = money[5] + 1
			continue
		}
		money[bills[i]] += 1

		change := bills[i] - 5

		//当需要找零时
		if change > 10 {
			if money[10] >= 1 && money[5] >= 1 {
				money[5] = money[5] - 1
				money[10] = money[10] - 1

			} else if money[5] >= 3 {
				money[5] = money[5] - 3

			} else {
				return false
			}
		} else if change < 10 {

			if money[5] >= 1 {
				money[5] = money[5] - 1
				
			} else {
				return false
			}

		}
	}

	return true
}

// @lc code=end
