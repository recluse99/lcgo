/*
 * @lc app=leetcode.cn id=1497 lang=golang
 *
 * [1497] 检查数组对是否可以被 k 整除
 */
package lcgo

// @lc code=start
func canArrange(arr []int, k int) bool {

	l := len(arr)

	res := make([]int, k)

	// 以下标大小为取余的数
	for i := 0; i < l; i++ {

		res[(arr[i]%k+k)%k]++ //arr[i]有可能为负数 。 还有一种解决办法是取余后 abs

	}

	if res[0]%2 != 0 {
		return false
	}

	// 余数为 i 和 k-i 的个数是要相等的

	// 优化   i<k/2
	for i := 1; i < k; i++ {
		if res[i] != res[k-i] {
			return false
		}
	}

	return true
}

// @lc code=end
