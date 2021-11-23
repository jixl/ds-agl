package structs

type TreeNode struct {
	Val         int
	next        *TreeNode
	Left, Right *TreeNode
}

func NewTree(val []interface{}) *TreeNode {
	size := len(val)
	if size == 0 {
		return nil
	}
	var nodes = make([]*TreeNode, size)
	for i, v := range val {
		if v == nil {
			nodes[i] = nil
		} else {
			nodes[i] = &TreeNode{Val: v.(int)}
		}
	}

	return queueInit(nodes)
}

func queueInit(nodes []*TreeNode) *TreeNode {
	queue, n := []*TreeNode{nodes[0]}, len(nodes)
	for i := 1; len(queue) > 0; {
		node := queue[0]
		queue = queue[1:]
		for j := 0; j < 2 && i < n; i, j = i+1, j+1 {
			if i%2 == 0 {
				node.Right = nodes[i]
			} else {
				node.Left = nodes[i]
			}
			if nodes[i] != nil {
				queue = append(queue, nodes[i])
			}
		}
	}
	return nodes[0]
}

// 先序遍历
func (root *TreeNode) PreOrder() []int {
	data := []int{}
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		data = append(data, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return data
}

// 中序遍历
func (root *TreeNode) InOrder() []int {
	data := []int{}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		data = append(data, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return data
}

// 后序遍历
func (root *TreeNode) PostOrder() []int {
	data := []int{}
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		data = append(data, node.Val)
	}
	postorder(root)
	return data
}

// 层序遍历
func (root *TreeNode) LevelOrder() []int {
	if root == nil {
		return nil
	}

	// res, que := [][]int{}, []*TreeNode{}
	res, que := []int{}, []*TreeNode{}
	que = append(que, root)
	for len(que) > 0 {
		n := len(que)
		// tmp := make([]int, 0)
		for i := 0; i < n; i++ {
			node := que[i]
			// tmp = append(tmp, node.Val)
			res = append(res, node.Val)
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		// res = append(res, tmp)
		que = que[n:]
	}
	return res
}
