package leetcode

// 876. 链表的中间结点
/**
 * Definition for singly-linked list.
 */
func middleNode(head *ListNode) *ListNode {
	mid, size := 0, 0

	for node := head; node != nil; node = node.Next {
		size++
	}

	if size%2 == 0 {
		mid = size / 2
	} else {
		mid = (size - 1) / 2
	}

	for i := 0; head != nil; head = head.Next {
		if i == mid {
			return head
		}
		i++
	}
	return nil
}
