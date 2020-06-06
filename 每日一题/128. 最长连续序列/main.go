package main

import "fmt"

// 给定一个未排序的整数数组，找出最长连续序列的长度。
//
//要求算法的时间复杂度为 O(n)。
//
//示例:
//
//输入: [100, 4, 200, 1, 3, 2]
//输出: 4
//解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-consecutive-sequence
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func longestConsecutive(nums []int) int {
	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}
	longestStreak := 0
	for num := range numSet {
		if _, ok := numSet[num-1]; !ok {
			currentNum := num
			currentStreak := 1
			for _, ok := numSet[currentNum+1]; ok; _, ok = numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
