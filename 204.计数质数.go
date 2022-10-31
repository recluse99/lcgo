/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */
package lcgo

// @lc code=start
func countPrimes(n int) int {
	if n < 2 {
		return 0
	}

	res := make([]int, n)
	var count = 0

	for i := 2; i < n; i++ {
		res[i] = 1
	}

	// 2 开始 。将2的倍数 去掉 ，下一个就是新的质数3.再
	//把3的倍数全部去掉。下一个又是新的倍数。5.。类推

	// 遍历到 根号i 就行
	for i := 2; i*i < n; i++ {
		if res[i] == 1 {
			count++
			// 去掉质数的倍数 。依次类推
			for j := i * 2; j < n; j = j + i {
				res[j] = 0
			}

		}

	}

	return count
}

// @lc code=end

//会超时

// func countPrimes(n int) int {

// 	if n < 2 {
// 		return 0
// 	}
// 	var count int = 0

// 	var flag = true

// 	for i := 2; i < n; i++ {
// 		// 重置  不要忘记
// 		flag = true
// 		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
// 			if i%j == 0 {
// 				flag = false
// 				break
// 			}
// 		}

// 		if flag {
// 			count++

// 		}
// 	}

// 	return count
// }
