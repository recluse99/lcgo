/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */
package lcgo

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var pre, temp, last *ListNode = head, head, nil

	for pre != nil {

		// pre 先行一步
		pre = pre.Next
		// 指针反指
		temp.Next = last

		// temp ,last 向前一位

		last = temp

		temp = pre

	}

	return last

}

// @lc code=end
