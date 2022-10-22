/*
 * @lc app=leetcode.cn id=717 lang=golang
 *
 * [717] 1 比特与 2 比特字符
 */
package lcgo

// @lc code=start
func isOneBitCharacter(bits []int) bool {

	l := len(bits)

	if l == 1 {
		return true
	}

	//最后一位是0
	// 遇到1就跳两步
	var index int
	for ; index < l-1; index++ {
		if bits[index] == 1 {
			index++ //
		}
	}
	//if  index停留在len(bits)-1上 代表l-3 和l-2构成了第二种字符，则最后一位一定是第一种字符 例 ...1x0
	//if  index  dt  l  上 ,说明 l-2上是1 则这时的l-2 和l-1构成的第二种字符的10情况 例  ...x10

	// if index == l-1 {
	// 	return true
	// } else if index == l {
	// 	return false
	// }

	// return false
	
	if index == l-1 {
		return true
	} else {
		return false
	}
}

// @lc code=end
