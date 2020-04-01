package main

import "fmt"

//有效括号字符串 定义：对于每个左括号，都能找到与之对应的右括号，反之亦然。详情参见题末「有效括号字符串」部分。
//
//嵌套深度 depth 定义：即有效括号字符串嵌套的层数。详情参见题末「嵌套深度」部分。
//
//
//
//给你一个「有效括号字符串」 seq，请你将其分成两个不相交的有效括号字符串 A 和 B：
//
//不相交：每个 seq[i] 只能分给 A 和 B 二者中的一个，不能既属于 A 也属于 B 。
//A 或 B 中的元素在原字符串中可以不连续。
//A.length + B.length = seq.length
//请你选出 任意 这样的 A 和 B，使得 max(depth(A), depth(B)) 的可能取值最小。其中 depth(A) 表示 A 的嵌套深度，depth(B) 表示 B 的嵌套深度。
//
//请你返回一个长度为 seq.length 的答案数组 answer，编码规则如下：
//
//answer[i] = 0，seq[i] 分给 A 。
//answer[i] = 1，seq[i] 分给 B 。
//即便有多个满足要求的答案存在，你也只需返回其中任意 一个 。
//
//
//
//示例 1：
//
//输入：seq = "(()())"
//输出：[0,1,1,1,1,0]
//示例 2：
//
//输入：seq = "()(())()"
//输出：[0,0,0,1,1,0,1,1]
//
//
//提示：
//
//1 <= text.size <= 10000
//
//
//有效括号字符串：
//
//仅由 "(" 和 ")" 构成的字符串，对于每个左括号，都能找到与之对应的右括号，反之亦然。
//
//下述几种情况同样属于有效括号字符串：
//
//空字符串
//连接，可以记作 AB（A 与 B 连接），其中 A 和 B 都是有效括号字符串
//嵌套，可以记作 (A)，其中 A 是有效括号字符串
//嵌套深度：
//
//类似地，我们可以定义任意有效括号字符串 s 的 嵌套深度 depth(S)：
//
//s 为空时，depth("") = 0
//s 为 A 与 B 连接时，depth(A + B) = max(depth(A), depth(B))，其中 A 和 B 都是有效括号字符串
//s 为嵌套情况，depth("(" + A + ")") = 1 + depth(A)，其中 A 是有效括号字符串
//例如：""，"()()"，和 "()(()())" 都是有效括号字符串，嵌套深度分别为 0，1，2，而 ")(" 和 "(()" 都不是有效括号字符串。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
// 题解：要使max(dept(A),dept(B))最小，则应尽量使A,B平分seq
// 由于有效字符串左右括号一定匹配，所以可以将左括号嵌套深度分为奇数层和偶数层，分别对应A和B即可
// 而求解嵌套深度，其实就是计算当前的栈深度
func maxDepthAfterSplit(seq string) []int {
	var ans []int
	var stack = -1
	for _, v := range seq {
		if v == '(' {
			stack++
			ans = append(ans, stack%2)
		} else if v == ')' {
			ans = append(ans, stack%2)
			stack--
		}
	}
	return ans
}

func main() {
	fmt.Println(maxDepthAfterSplit("()(())()"))
}
