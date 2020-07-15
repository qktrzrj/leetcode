package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		l1, l2 int
		r1, r2 = len(nums1) - 1, len(nums2) - 1
	)
	find := func() (min, max int) {
		min, max = math.MaxInt64, math.MinInt64
		var n1, n2 = l1 <= r1, l2 <= r2
		if n1 {
			min, max = nums1[l1], nums1[r1]
			l1++
			r1--
		}
		if n2 {
			if nums2[l2] <= min {
				min = nums2[l2]
				l2++
				if n1 {
					l1--
				}
			}
			if nums2[r2] >= max {
				max = nums2[r2]
				r2--
				if n1 {
					r1++
				}
			}
		}
		return
	}
	var l, r int
	for l1 <= r1 || l2 <= r2 {
		l, r = find()
	}
	return (float64(l) + float64(r)) / 2
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2, 4}))
}
