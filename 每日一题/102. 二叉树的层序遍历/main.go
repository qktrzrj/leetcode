package main

import "fmt"

// 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
//
//
//示例：
//二叉树：[3,9,20,null,null,15,7],
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//返回其层次遍历结果：
//
//[
//  [3],
//  [9,20],
//  [15,7]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	gap := &TreeNode{}
	stack := []*TreeNode{root, gap}
	row := make([]int, 0)
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		if node == gap {
			ans = append(ans, row)
			row = make([]int, 0)
			if len(stack) > 0 {
				stack = append(stack, gap)
			}
			continue
		}
		row = append(row, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return ans
}

func main() {
	fmt.Println(levelOrder(nil))
}

func NewTree(s []int) *TreeNode {
	root := &TreeNode{Val: s[0]}
	stack := []*TreeNode{root}
	i := 1
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
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
	return root
}
