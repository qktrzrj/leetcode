package main

import "fmt"

//给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。
//
//说明：
//你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。
//
//示例 1:
//
//输入: root = [3,1,4,null,2], k = 1
//   3
//  / \
// 1   4
//  \
//   2
//输出: 1
//示例 2:
//
//输入: root = [5,3,6,2,4,null,null,1], k = 3
//       5
//      / \
//     3   6
//    / \
//   2   4
//  /
// 1
//输出: 3
//进阶：
//如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
// BST的遍历，由于BST的特点，节点K的值一定大于左子树中节点值，小于右子树中节点值
// 树的遍历方法分为前序遍历，中序遍历，后序遍历
// BST的中序遍历，结果为生序序列，可以直接求得第K大的值

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	for {
		//遍历左子树，按倒序入栈，最后一个节点即为最小节点
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(kthSmallest(root, 2))
}
