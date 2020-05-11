package main

import "fmt"

// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
//百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
//
//例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,-1,-1,7,4]
//
//
//
//
//
//示例 1:
//
//输入: root = [3,5,1,6,2,0,8,-1,-1,7,4], p = 5, q = 1
//输出: 3
//解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
//示例 2:
//
//输入: root = [3,5,1,6,2,0,8,-1,-1,7,4], p = 5, q = 4
//输出: 5
//解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
//
//
//说明:
//
//所有节点的值都是唯一的。
//p、q 为不同节点且均存在于给定的二叉树中。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// Definition for TreeNode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	lChild := lowestCommonAncestor(root.Left, p, q)
	rChild := lowestCommonAncestor(root.Right, p, q)
	if lChild == nil {
		return rChild
	} else if rChild == nil {
		return lChild
	} else {
		return root
	}
}

func main() {
	root, p, q := NewTree([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}, 5, 1)
	fmt.Println(lowestCommonAncestor(root, p, q))
}

func NewTree(s []int, p, q int) (*TreeNode, *TreeNode, *TreeNode) {
	root := &TreeNode{Val: s[0]}
	stack := []*TreeNode{root}
	i := 1
	var pn, qn *TreeNode
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		if node.Val == p {
			pn = node
		}
		if node.Val == q {
			qn = node
		}
		if i < len(s) && s[i] != -1 {
			node.Left = &TreeNode{Val: s[i]}
			stack = append(stack, node.Left)
		}
		if i+1 < len(s) && s[i+1] != -1 {
			node.Right = &TreeNode{Val: s[i+1]}
			stack = append(stack, node.Right)
		}
		i += 2
	}
	return root, pn, qn
}
