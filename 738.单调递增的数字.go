/*
 * @lc app=leetcode.cn id=738 lang=golang
 *
 * [738] 单调递增的数字
 */
package lcgo

import (
	"strconv"
)

// @lc code=start
func monotoneIncreasingDigits(n int) int {
	
	if n==0{
		return 0
	}

	str := strconv.FormatInt(int64(n), 10)

	numbers := make([]int, len(str))

	// string 的单字符类型为int32 等价于rune 汉字占3个字节
	for i, value := range str {

		numbers[i] = int(value) - 48

	}


	//从后往前遍历
	for i := len(numbers) - 1; i > 0; i-- {
		if numbers[i-1] > numbers[i] {
			numbers[i-1] -= 1

			//将index后面的数字全部改为9
			for j := i; j < len(numbers); j++ {
				numbers[j] = 9
			}
		}

	}

	var result int

	//将字符串转为数字
	for i := 0; i < len(numbers); i++ {

		result=result*10+numbers[i]
	}

	return result

}

// @lc code=end
