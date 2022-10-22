/*
 * @lc app=leetcode.cn id=697 lang=golang
 *
 * [697] 数组的度
 */
package lcgo

// @lc code=start
func findShortestSubArray(nums []int) int {

	l := len(nums)

	if l == 1 {
		return 1
	}
	//最小的数组长度
	var smallLen = l
	// map 不用直接开辟很大的空间
	resultMap := map[int][]int{}
	//数字出现的长度
	maxCount := 0

	for i, value := range nums {
		//map的key是否存在value
		if _, ok := resultMap[value]; !ok {
			//[]int{} 中 0：出现的次数  1：第一次出现的索引 2：最后一次出现的索引
			resultMap[value] = []int{1, i, i}
		} else {
			//频数加一
			resultMap[value][0]++
			//索引更新
			resultMap[value][2] = i
		}

		if resultMap[value][0] > maxCount {
			maxCount = resultMap[value][0]
		}
	}

	for _, temp := range resultMap {
		if temp[0] == maxCount {
			if temp[2]-temp[1]+1 < smallLen {
				smallLen = temp[2] - temp[1] + 1
			}
		}
	}

	return smallLen
}

// @lc code=end
