/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */
package lcgo

// @lc code=start

var dMap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}
var result []string

var temp string

func letterCombinations(digits string) []string {

	l := len(digits)

	if l == 0 {
		return []string{}
	}

	// 初始化  ，清空全局变量 。不然在leetcode中提交中可能会1报错
	result = []string{}

	return recur(digits, 0)
}

func recur(digits string, index int) []string {

	if index == len(digits) {
		result = append(result, temp)

		return nil
	}

	res := dMap[digits[index]]

	for j := 0; j < len(res); j++ {
		temp = temp + string(res[j])
		recur(digits, index+1)

		// 回溯
		temp = temp[:len(temp)-1]

	}

	//temp = temp[:len(temp)-1]

	return result
}

// @lc code=end
