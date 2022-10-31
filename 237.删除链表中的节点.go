/*
 * @lc app=leetcode.cn id=237 lang=golang
 *
 * [237] 删除链表中的节点
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
func deleteNode(node *ListNode) {

	// 替身攻击  hahaha  
	// 挺好玩的。。。

	node.Val = node.Next.Val

	node.Next = node.Next.Next

}

// @lc code=end
