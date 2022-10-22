/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] 种花问题
 */
package lcgo

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {

	if n == 0 {
		return true
	}

	//防御性种花
	// 在 花坛两侧加上0  ，这时 ，只要出现三个连续的0 ，就可以种上一朵花

	a := []int{0}
	flowerbed = append(a, flowerbed...)
	flowerbed = append(flowerbed, a...)
	l := len(flowerbed)
	for i := 0; i < l-2; i++ {

		if flowerbed[i] == 0 && flowerbed[i+1] == 0 && flowerbed[i+2] == 0 {
			n--
			// 比种花的index 大一个   
			i = i + 1   // 因为i还会执行i++ ,所以只用加一
		}

	}

	if n <= 0 {
		return true
	} else {
		return false
	}

}

// @lc code=end
