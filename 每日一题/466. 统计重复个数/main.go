package main

import "fmt"

//由 n 个连接的字符串 s 组成字符串 S，记作 S = [s,n]。例如，["abc",3]=“abcabcabc”。
//
//如果我们可以从 s2 中删除某些字符使其变为 s1，则称字符串 s1 可以从字符串 s2 获得。
// 例如，根据定义，"abc" 可以从 “abdbec” 获得，但不能从 “acbbe” 获得。
//
//现在给你两个非空字符串 s1 和 s2（每个最多 100 个字符长）和两个整数 0 ≤ n1 ≤ 106 和 1 ≤ n2 ≤ 106。
// 现在考虑字符串 S1 和 S2，其中 S1=[s1,n1] 、S2=[s2,n2] 。
//
//请你找出一个可以满足使[S2,M] 从 S1 获得的最大整数 M 。
//
//
//
//示例：
//
//输入：
//s1 ="acb",n1 = 4
//s2 ="ab",n2 = 2
//
//返回：
//2
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/count-the-repetitions
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
// 看题意，可知为，求一个整数M，使[s2,n2*M]可以从[s1,n1]中获得，要求M最大。
// 既然要M最大，意味着从S1中删除的字符最少，即找到在S1中，s2可获取的最大数量。
// 假设S1=acbacbacbacb,s2=ab,我们将c全部删除，得到abababab,可知s2出现4次，M=4/n2=2
// 那么最后问题变成，在S1中找出s2的次数。
// 我们对字符串进行匹配，当遇到不匹配的字符时，处理方法只能为删除
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	var S1 string
	for i := 0; i < n1; i++ {
		S1 += s1
	}
	var ans int
	for i := 0; i < len(S1); {
		for j := 0; j < len(s2); {
			if i > len(S1)-1 {
				goto Return
			}
			if S1[i] == s2[j] {
				i++
				j++
				continue
			}
			i++
		}
		ans++
	}
Return:
	return ans / n2
}

func main() {
	fmt.Println(getMaxRepetitions("acb", 2, "abc", 1))
}
