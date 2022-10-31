/*
 * @lc app=leetcode.cn id=941 lang=golang
 *
 * [941] 有效的山脉数组
 */
package lcgo

// @lc code=start
func validMountainArray(arr []int) bool {

	// 双指针法

	l := len(arr)

	if l < 3 {
		return false
	}
	
	var isFall = false

	for i := 0; i < l-1; i++ {

		// 开始就fall  return false
		if i == 0 && arr[i] >= arr[i+1] {
			return false
		}

		   // up
		if !isFall {
			if arr[i] > arr[i+1] {
				isFall = true

			} else if arr[i] == arr[i+1] {
				return false
			}
			// fall
		} else {
			if arr[i] <= arr[i+1] {
				return false
			}
		}
	}

	// 存在fall
	if !isFall {
		return false
	}

	return true
}

// @lc code=end



/**
func validMountainArray(arr []int) bool {

	l := len(arr)

	if l < 3 {
		return false
	}

	var isFall = false

	for i := 0; i < l-1; i++ {

		// 开始就fall  return false
		if i == 0 && arr[i] >= arr[i+1] {
			return false
		}

		   // up
		if !isFall {
			if arr[i] > arr[i+1] {
				isFall = true

			} else if arr[i] == arr[i+1] {
				return false
			}
			// fall
		} else {
			if arr[i] <= arr[i+1] {
				return false
			}
		}
	}

	// 存在fall
	if !isFall {
		return false
	}

	return true
}

**/