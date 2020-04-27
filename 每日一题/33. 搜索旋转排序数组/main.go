package main

import "fmt"

// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//
//( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
//
//搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
//
//你可以假设数组中不存在重复的元素。
//
//你的算法时间复杂度必须是 O(log n) 级别。
//
//示例 1:
//
//输入: nums = [4,5,6,7,0,1,2], target = 0
//输出: 4
//示例 2:
//
//输入: nums = [4,5,6,7,0,1,2], target = 3
//输出: -1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)
		if nums[h] < nums[0] {
			// 旋转点位于i-h之间
			j = h
		} else {
			i = h + 1
		}
	}
	// 没有旋转
	if i >= n {
		i, j = 0, n
	} else if nums[0] <= target { // 旋转点为i,搜索前半部分
		i, j = 0, i-1
	} else { // 旋转点为i,搜索后半部分
		j = n
	}

	for i < j {
		h := int(uint(i+j) >> 1)
		if nums[h] < target {
			i = h + 1
		} else {
			j = h
		}
	}
	if i >= n {
		return -1
	}
	if nums[i] == target {
		return i
	}
	return -1
}

func main() {
	fmt.Println(search([]int{3, 1}, 3))
}
