package lcgo

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	Map := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		//判断是否有key存在
		index, ok := Map[target-nums[i]]
		if ok {
			return []int{index, i}
		} else {
			Map[nums[i]] = i
		}
	}

	return nil
}

// @lc code=end
