package main

import "fmt"

// 给定一个整数数组 nums，将该数组升序排列。
//
// 示例 1：
// 输入：[5,2,3,1]
// 输出：[1,2,3,5]
//
// 示例 2：
// 输入：[5,1,1,2,0,0]
// 输出：[0,0,1,1,2,5]
func sortArray(nums []int) []int {
	var quickSort func(start, end int)
	quickSort = func(start, end int) {
		if start >= end {
			return
		}
		pivot := nums[end]
		p := start
		for q := p; q < end; q++ {
			if nums[q] < pivot {
				nums[p], nums[q] = nums[q], nums[p]
				p++
			}
		}
		nums[p], nums[end] = nums[end], nums[p]
		quickSort(start, p-1)
		quickSort(p+1, end)
	}
	quickSort(0, len(nums)-1)
	return nums
}

func main() {
	fmt.Println(sortArray([]int{5, 2, 3, 1}))
}
