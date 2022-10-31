/*
 * @lc app=leetcode.cn id=914 lang=golang
 *
 * [914] 卡牌分组
 */
package lcgo

// @lc code=start
func hasGroupsSizeX(deck []int) bool {

	resultMap := map[int]int{}
	//最大公因数初始值
	g := -1
	for _, d := range deck {
		resultMap[d]++
	}

	for _, count := range resultMap {
		if count < 2 {
			return false
		}

		if g == -1 {
			g = count
		} else {
			g = gcd(g, count)
		}

	}

	return g >= 2
}

// 求两个数的最大公因数
func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

// @lc code=end
