/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */
package lcgo

// 防止报错  添加包结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil && list2 == nil {
		return nil
	}

	result := &ListNode{}

	// 临时指针
	temp := result

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			temp.Next = &ListNode{list1.Val, nil}
			list1 = list1.Next
			temp = temp.Next

		} else {
			temp.Next = &ListNode{list2.Val, nil}
			list2 = list2.Next
			temp = temp.Next
		}

	}

	if list1 == nil {
		temp.Next = list2
	} else {
		temp.Next = list1
	}

	return result.Next

}

// @lc code=end
