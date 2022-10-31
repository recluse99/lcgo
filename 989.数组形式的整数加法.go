/*
 * @lc app=leetcode.cn id=989 lang=golang
 *
 * [989] 数组形式的整数加法
 */
package lcgo

// @lc code=start
func addToArrayForm(num []int, k int) []int {

	result := []int{}

	l := len(num)
	var a int

	for i := l - 1; i >= 0; i-- {
		sum := num[i] + k%10 + a

		//k 去掉尾
		k /= 10

		if sum >= 10 {
			a = 1
			sum -= 10
		} else {
			a = 0
		}
		result = append(result, sum)
	}

	if a != 0 {
		k++
	}

	// 如果k还剩下

	for ; k > 0; k = k / 10 {
		result = append(result, k%10)
	}

	// reverse result

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// @lc code=end

/*
func addToArrayForm(num []int, k int) []int {

	arr := getArrayByK(k)

	lArr := len(arr)
	num = append([]int{0}, num...)
	lNum := len(num)
	var a int
	var temp int
	result := make([]int, 0, lNum+1)

	for i := 0; i < lArr; i++ {
		if a == 1 {
			if nums := arr[lArr-1-i] + num[lNum-1-i] + 1; nums > 10 {
				a = 1
				temp = nums - 10
			} else {
				temp = nums
				a = 0
			}
			//无进位
		} else {
			if nums := arr[lArr-1-i] + num[lNum-1-i]; nums > 10 {
				a = 1
				temp = nums - 10
			} else {
				temp = nums
				a = 0
			}

		}

		//
		result = append(result, temp)

	}

	//处理 num 剩下的部分
	for i := 0; i < lNum-lArr; i++ {
		result = append(result, num[lNum-lArr-1-i])
	}

	// 逆转切片
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {

		result[i], result[j] = result[j], result[i]

	}

	return result

}

// 将 int k conv 为 []int
func getArrayByK(k int) []int {

	str := strconv.Itoa(k)

	l := len(str)

	result := make([]int, l+1)

	for i, v := range str {

		result[i+1] = int(v) - 48

	}

	return result
}




*/
