package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(val []int) *ListNode {
	var list, node *ListNode
	for i, v := range val {
		tmp := &ListNode{Val: v}
		if i == 0 {
			list, node = tmp, tmp
		} else {
			node.Next = tmp
			node = node.Next
		}
	}
	return list
}

func (list *ListNode) toArray() []int {
	var arr []int
	for list != nil {
		arr = append(arr, list.Val)
		list = list.Next
	}
	return arr
}
