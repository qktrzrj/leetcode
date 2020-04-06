package main

import "fmt"

//给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
//
//你可以对一个单词进行如下三种操作：
//
//插入一个字符
//删除一个字符
//替换一个字符
//
//
//示例 1：
//
//输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
//示例 2：
//
//输入：word1 = "intention", word2 = "execution"
//输出：5
//解释：
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/edit-distance
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
// 动态规划解法:
// 我们对两个字符串从左到右进行字符比对，当字符不相同时，可以进行插入，删除，替换三个操作以达到使字符相同的目的
// 但是实际上我们不需要实际执行这些操作
// 设字符串为a，b, 求min(dis(a,b))
// 对于插入，代表当前a字符索引不变，b字符认为匹配成功索引向后移动一位
// 对于替换，直接将a和b字符索引均移动一位
// 对于删除，代表a字符索引向后一位，重新匹配
// 则对于a[i],b[j],若a[i]！=b[j],则可以选择插入即=dis(a[i],b[j+1])+1,删除=dis(a[i+1],b[j])+1,替换=dis(a[i+1],b[j+1])+1
// 若a[i]==b[j],则dis(a[i+1],b[j+1])
// 可以观察到，若a[i]==b[j],那么可能是因为正常匹配或替换后匹配成功（对于替换后成功匹配，由于i-1,j-1肯定不匹配，
// 距离可在前一步计算） 即 = dis(a[i-1],b[j-1])
// 也可能是插入后匹配 = dis(a[i],b[j-1])+1,删除后匹配 = dis(a[i-1],b[j])+1
// 如果a[i]!=b[j],那么可能是因为正常匹配或替换后匹配不成功(默认进行替换) = dis(a[i-1],b[j-1])+1
// 也可能是插入后不匹配 = dis(a[i],b[j-1])+1,删除后不匹配 = dis(a[i-1],b[j])+1
// 因为涉及i-1和j-1，所以我们的迭代应该在到达边界时结束
func min(i, j, k int) int {
	if i < j {
		if i < k {
			return i
		}
		return k
	} else if i < k {
		return j
	} else if k < j {
		return k
	}
	return j
}

func minDistance(word1 string, word2 string) int {
	a, b := []byte(word1), []byte(word2)
	n, m := len(a), len(b)
	if n*m == 0 {
		return n + m
	}
	minDist := make([][]int, n+1)
	for index := 0; index < n+1; index++ {
		minDist[index] = make([]int, m+1)
		minDist[index][0] = index
	}
	for index := 0; index < m+1; index++ {
		minDist[0][index] = index
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if a[i-1] == b[j-1] {
				minDist[i][j] = min(minDist[i-1][j]+1, minDist[i][j-1]+1, minDist[i-1][j-1])
			} else {
				minDist[i][j] = min(minDist[i-1][j]+1, minDist[i][j-1]+1, minDist[i-1][j-1]+1)
			}
		}
	}
	return minDist[n][m]
}

func main() {
	fmt.Println(minDistance("dinitrophenylhydrazine", "acetylphenylhydrazine"))
}
