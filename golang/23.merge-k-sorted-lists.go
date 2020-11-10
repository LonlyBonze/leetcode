/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//1
// func mergeKLists(lists []*ListNode) *ListNode {
// 	if len(lists) == 0 {
// 		return nil
// 	}
// 	if len(lists) == 1 {
// 		return lists[0]
// 	}

// 	head := &ListNode{}
// 	tail := head
// 	for {
// 		if lists[0] == nil {
// 			tail.Next = lists[1]
// 			break
// 		}
// 		if lists[1] == nil {
// 			tail.Next = lists[0]
// 			break
// 		}
// 		if lists[0].Val < lists[1].Val {
// 			tail.Next = lists[0]
// 			lists[0] = lists[0].Next
// 		} else {
// 			tail.Next = lists[1]
// 			lists[1] = lists[1].Next
// 		}
// 		tail = tail.Next
// 	}
// 	lists[1] = head.Next

// 	return mergeKLists(lists[1:])
// }

//2
func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{}
	tail := head

	for i := 0; i < len(lists); {
		if lists[i] == nil {
			if len(lists) == 1 {
				return nil
			}
			lists = append(lists[:i], lists[i+1:]...)
			continue
		}
		i++
	}

	for i := 0; i < len(lists)-1; i++ {
		for j := len(lists) - 1; j > i; j-- {
			if lists[j-1].Val > lists[j].Val {
				lists[j-1], lists[j] = lists[j], lists[j-1]
			}
		}
	}

	for len(lists) > 0 {
		if len(lists) == 1 {
			tail.Next = lists[0]
			break
		}

		tail.Next = lists[0]
		tail = lists[0]

		if lists[0].Next == nil {
			lists = lists[1:]
			continue
		}

		lists[0] = lists[0].Next
		for i := 0; i < len(lists)-1; i++ {
			if lists[i].Val > lists[i+1].Val {
				lists[i], lists[i+1] = lists[i+1], lists[i]
			} else {
				break
			}
		}
	}

	return head.Next
}

// @lc code=end

