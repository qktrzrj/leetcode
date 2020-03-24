package test_1052_test

import (
	"fmt"
	"testing"
)

// 今天，书店老板有一家店打算试营业 customers.length 分钟。
// 每分钟都有一些顾客（customers[i]）会进入书店，所有这些顾客都会在那一分钟结束后离开。
//
//在某些时候，书店老板会生气。 如果书店老板在第 i 分钟生气，那么 grumpy[i] = 1，否则 grumpy[i] = 0。
// 当书店老板生气时，那一分钟的顾客就会不满意，不生气则他们是满意的。
//
//书店老板知道一个秘密技巧，能抑制自己的情绪，可以让自己连续 X 分钟不生气，但却只能使用一次。
//
//请你返回这一天营业下来，最多有多少客户能够感到满意的数量。
//
//
//示例：
//
//输入：customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], X = 3
//输出：16
//解释：
//书店老板在最后 3 分钟保持冷静。
//感到满意的最大客户数量 = 1 + 1 + 1 + 1 + 7 + 5 = 16.
//
//
//提示：
//
//1 <= X <= customers.length == grumpy.length <= 20000
//0 <= customers[i] <= 1000
//0 <= grumpy[i] <= 1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/grumpy-bookstore-owner
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
// 解题思路：用滑动窗口解题
// 这个题，可以理解为，哪个连续时间为X的时间段内，不满意的客户数量最多
// 那么我们可以设一个滑动窗口,窗口长度为X，初始时，窗口起始位置为数据组起始位置，然后向右滑动窗口，计算滑动前后的不满意客户数，取最大值
// 此题的难点再于，若要先初始化好一个滑动窗口，则必须先遍历数组的0...X段，然后再对数组全遍历
// 其实我们可以将窗口设为一个可增长窗口，并设定最大长度为X，当窗口长度增长超过X时，向左收缩窗口边界，以此实现窗口的滑动
func maxSatisfied(customers []int, grumpy []int, X int) int {
	var left, sum, dissatisfied, maxDiss int
	for right, num := range customers {
		if grumpy[right] == 0 {
			sum += num
		} else {
			dissatisfied += num
		}
		if right-left+1 == X {
			maxDiss = max(maxDiss, dissatisfied)
			if grumpy[left] == 1 {
				dissatisfied -= customers[left]
			}
			left++
		}
	}
	return sum + max(maxDiss, dissatisfied)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Benchmark_MaxSatisfied(b *testing.B) {
	b.ReportAllocs()
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
}
