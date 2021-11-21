package leetcode

// import "fmt"

// 116. 填充每个节点的下一个右侧节点指针 https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
}

func connect(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	node := root
	clr(node.Left, node.Right)
	return node
}

func clr(left, right *TreeNode) {
	if left != nil {
		left.Next = right
		lr, rl := left.Right, right.Left
		for lr != nil {
			lr.Next = rl
			lr, rl = lr.Right, rl.Left
		}
		clr(left.Left, left.Right)
		clr(right.Left, right.Right)
	}
}
