package main

import (
	"fmt"
	"sort"
)

// 给你一个整数数组 arr 和一个目标值 target ，请你返回一个整数 value ，使得将数组中所有大于 value 的值变成 value 后，数组的和最接近  target （最接近表示两者之差的绝对值最小）。
//
//如果有多种使得和最接近 target 的方案，请你返回这些整数中的最小值。
//
//请注意，答案不一定是 arr 中的数字。
//
//
//
//示例 1：
//
//输入：arr = [4,9,3], target = 10
//输出：3
//解释：当选择 value 为 3 时，数组会变成 [3, 3, 3]，和为 9 ，这是最接近 target 的方案。
//示例 2：
//
//输入：arr = [2,3,5], target = 10
//输出：5
//示例 3：
//
//输入：arr = [60864,25176,27249,21296,20204], target = 56803
//输出：11361
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sum-of-mutated-array-closest-to-target
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
	}
	l, r, ans := 0, arr[n-1], -1
	for l <= r {
		mid := (l + r) >> 1
		index := sort.SearchInts(arr, mid)
		if index < 0 {
			index = 0
		}
		cur := prefix[index] + (n-index)*mid
		if cur <= target {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	chooseSmall := check(arr, ans)
	chooseBig := check(arr, ans+1)
	if abs(chooseSmall-target) > abs(chooseBig-target) {
		ans++
	}
	return ans
}

func check(arr []int, x int) int {
	ret := 0
	for _, num := range arr {
		ret += min(num, x)
	}
	return ret
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func main() {
	fmt.Println(findBestValue([]int{3, 9, 1}, 10))
}
