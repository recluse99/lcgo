/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */
package lcgo

import "fmt"

// @lc code=start
func addBinary(a string, b string) string {

	al := len(a)
	bl := len(b)
	var result string
	if al >= bl {
		result = addBin(a, b)
	} else {
		result = addBin(b, a)
	}

	return result
}

// l:a  > l:b
func addBin(a string, b string) string {

	la := len(a)
	lb := len(b)
	//进位
	var c int
	var result string
	for i := 0; i < lb; i++ {

		temp := int(a[la-1-i]-48+b[lb-1-i]-48) + c
		if temp == 2 {
			c = 1
			temp = 0
		} else {
			c = 0
		}

		result = fmt.Sprintf("%d%s", temp, result)

	}
	// 处理剩下的字符串

	for i :=0; i < la-lb; i++ {
		temp := int(a[la-lb-1-i]-48) + c
		if temp == 2 {
			c = 1
			temp = 0
		} else {
			c = 0
		}
		result = fmt.Sprintf("%d%s", temp, result)
	}
	if c != 0 {
		result = fmt.Sprintf("1%s", result)
	}

	return result
}

// @lc code=end
