/*
 * @lc app=leetcode.cn id=888 lang=golang
 *
 * [888] 公平的糖果交换
 */
package lcgo

// @lc code=start
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {

	bobMap := map[int]int{}

	aliceSum := sum888(aliceSizes)

	var bobSum int

	for index, value := range bobSizes {

		// 主要需要value  index覆盖了无所谓
		bobMap[value] = index
		bobSum += value
	}

	//bobsum-bob+alice=alicesum-alice+bob

	for _, value := range aliceSizes {

		bobGive := (bobSum - aliceSum + 2*value) / 2
		if _, ok := bobMap[bobGive]; ok {
			return []int{value, bobGive}
		}

	}

	return []int{}
}

func sum888(a []int) int {
	var sum int

	for _, v := range a {
		sum += v
	}

	return sum
}

// @lc code=end
