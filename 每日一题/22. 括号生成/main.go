package main

import "fmt"

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
//示例：
//
//输入：n = 3
//输出：[
//       "((()))",
//       "(()())",
//       "(())()",
//       "()(())",
//       "()()()"
//     ]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/generate-parentheses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	var ans []string
	var loop func(p string, l, left, right int)
	loop = func(p string, l, left, right int) {
		if l > 0 {
			// choose (
			loop(p+"(", l-1, left+1, right)
			// choose )
			if right+1 == left {
				loop(p+")(", l-1, left+1, right+1)
			} else {
				loop(p+")", l, left, right+1)
			}
		} else {
			for i := 0; i < n-right; i++ {
				p += ")"
			}
			ans = append(ans, p)
		}
	}
	loop("(", n-1, 1, 0)
	return ans
}

func main() {
	fmt.Println(generateParenthesis(3))
}
