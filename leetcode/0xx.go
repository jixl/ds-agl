package leetcode

// import "fmt"

// 3. 无重复字符的最长子串 https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	max, mp := 1, make(map[byte]int)
	for l, r := 0, 0; r < n; r++ {
		if i, ok := mp[s[r]]; ok {
			for ; l <= i; l++ {
				delete(mp, s[l])
			}
		} else if max < r-l+1 {
			max = r - l + 1
		}
		mp[s[r]] = r
	}
	return max
}

// 21. 合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var list, node *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			list, node = _node(list, node, l2)
			l2 = l2.Next
			if l2 == nil {
				node.Next = l1
			}
		} else {
			list, node = _node(list, node, l1)
			l1 = l1.Next
			if l1 == nil {
				node.Next = l2
			}
		}
	}
	return list
}

func _node(list, node, l *ListNode) (*ListNode, *ListNode) {
	if list == nil {
		list = l
		node = l
	} else {
		node.Next = l
		node = node.Next
	}
	return list, node
}

// 53. 最大子序和
// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

// 19. 删除链表的倒数第 N 个结点
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var size int
	for node := head; node != nil; node = node.Next {
		size++
	}

	r := size - n
	for node, i := head, 1; node != nil; node = node.Next {
		if r == 0 {
			return head.Next
		} else if r == i {
			node.Next = node.Next.Next
			return head
		}
		i++
	}
	return nil
}

// 36. 有效的数独 https://leetcode-cn.com/problems/valid-sudoku
// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
// 输入：board =
// [["5","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// 输出：true
// 输入：board =
// [["8","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// 输出：false
func isValidSudoku(board [][]byte) bool {
	// 方格map
	mp := make(map[byte]bool)
	var i, j byte
	for i = 0; i < 9; i++ {
		var row, col byte
		// 行列map 行 j, 列 c+j
		lmp := make(map[byte]bool)
		for j = 0; j < 9; j++ {
			row, col = board[i][j], board[j][i]
			if row != '.' {
				if lmp[row] || mp[key(i, j, row)] {
					return false
				}
				lmp[row], mp[key(i, j, row)] = true, true
			}
			if col != '.' {
				if lmp['c'+col] {
					return false
				}
				lmp['c'+col] = true
			}
		}
	}
	return true
}

func key(i, j, v byte) byte {
	var x, y byte
	x, y = (i+3)/3, (j+3)/3
	return x*y + 10 + v
}

// 73. 矩阵置零 https://leetcode-cn.com/problems/set-matrix-zeroes
// 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
// 进阶：
// 一个直观的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
// 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
// 你能想出一个仅使用常量空间的解决方案吗？
// 示例 1：
// 输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
// 输出：[[1,0,1],[0,0,0],[1,0,1]]
// 示例 2：
// 输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// 输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]
func setZeroes(matrix [][]int) {
	y := len(matrix[0])
	var idx = make([]int, y)

	for _, mat := range matrix {
		z := 0
		for j := 0; j < y; j++ {
			if mat[j] == 0 {
				idx[j] = 1
				z = 1
			}
		}
		if z == 1 {
			for j := 0; j < y; j++ {
				mat[j] = 0
			}
		}
	}

	for j, v := range idx {
		if v == 1 {
			for _, mat := range matrix {
				mat[j] = 0
			}
		}
	}
}
