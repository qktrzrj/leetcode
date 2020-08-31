package main

import "fmt"

func maxArea(height []int) int {
	var left, right, max int
	right = len(height) - 1
	for left < right {
		tmp := (right - left) * minHeight(height[left], height[right])
		if tmp > max {
			max = tmp
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

func minHeight(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
