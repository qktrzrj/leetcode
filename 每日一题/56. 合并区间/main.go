package main

import (
	"fmt"
	"sort"
)

//给出一个区间的集合，请合并所有重叠的区间。
//
//示例 1:
//
//输入: [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//示例 2:
//
//输入: [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-intervals
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ss [][]int

func (s ss) Len() int {
	return len(s)
}

func (s ss) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func (s ss) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	is := ss(intervals)
	sort.Sort(is)
	ans := [][]int{is[0]}
	for i := 1; i < len(is); i++ {
		n := len(ans) - 1
		x, y, px, py := is[i][0], is[i][1], ans[n][0], ans[n][1]
		if x <= py {
			if y >= py {
				ans[n] = []int{px, y}
			}
		} else {
			ans = append(ans, is[i])
		}
	}
	return ans
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
