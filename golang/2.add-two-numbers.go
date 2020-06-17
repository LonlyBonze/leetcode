/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	head := l1
	for {
		l1.Val = l1.Val + l2.Val + carry
		carry =  l1.Val / 10
		l1.Val = l1.Val % 10
		if l1.Next == nil{
			if carry != 0 {
				if l2.Next == nil {
					l1.Next = &ListNode {
						Val: carry,
					}
				} else {
					l1.Next = &ListNode {}
					l1 = l1.Next
					l2 = l2.Next
					continue
				}
			} else {
				l1.Next = l2.Next
			}
			break
		} else if l2.Next == nil {
			if carry != 0 {
				if l1.Next == nil {
					l1.Next = &ListNode {
						Val: carry,
					}
				} else {
					l2.Next = &ListNode {}
					l1 = l1.Next
					l2 = l2.Next
					continue
				}
			}
			break
		} else {
			l1 = l1.Next
			l2 = l2.Next
		}
	}
	return head
}
// @lc code=end

